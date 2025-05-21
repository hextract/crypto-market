import React, { useEffect, useState } from 'react';
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  Legend,
  ResponsiveContainer,
  ReferenceLine,
  Label,
  ReferenceDot
} from 'recharts';
import { getCurvesData, getClearingPrice } from '../api/matchingEngineService';
import { useTranslation } from 'react-i18next';
import useInterval from '../hooks/useInterval';

const CustomizedAxisTick = ({ x, y, payload }) => {
  return (
    <text
      x={x}
      y={y}
      dy={16}
      textAnchor="middle"
      fill="#ffffff"
      fontSize={14}
      fontWeight={500}
    >
      {payload.value.toFixed(3)}
    </text>
  );
};

export default function MarketChart() {
  const { t } = useTranslation();
  const [supply, setSupply] = useState([]);
  const [demand, setDemand] = useState([]);
  const [intersection, setIntersection] = useState({ price: 0, volume: 0 });
  const [loading, setLoading] = useState(true);
  const [bounds, setBounds] = useState({ min: 0, max: 0 });

  // Функция для нахождения точки пересечения кривых
  const findIntersection = (supplyCurve, demandCurve, clearingPrice) => {
    // Сортируем кривые по цене
    const sortedSupply = [...supplyCurve].sort((a, b) => a.price - b.price);
    const sortedDemand = [...demandCurve].sort((a, b) => a.price - b.price);

    // Находим ближайшие точки вокруг цены клиринга
    let supplyBefore = 0, supplyAfter = 0;
    for (let i = 0; i < sortedSupply.length - 1; i++) {
      if (sortedSupply[i].price <= clearingPrice && sortedSupply[i+1].price >= clearingPrice) {
        supplyBefore = sortedSupply[i];
        supplyAfter = sortedSupply[i+1];
        break;
      }
    }

    let demandBefore = 0, demandAfter = 0;
    for (let i = 0; i < sortedDemand.length - 1; i++) {
      if (sortedDemand[i].price <= clearingPrice && sortedDemand[i+1].price >= clearingPrice) {
        demandBefore = sortedDemand[i];
        demandAfter = sortedDemand[i+1];
        break;
      }
    }

    // Если не нашли подходящих точек, используем крайние значения
    if (!supplyBefore || !demandBefore) {
      return {
        price: clearingPrice,
        volume: (sortedSupply[0].volume + sortedDemand[0].volume) / 2
      };
    }

    // Линейная интерполяция для нахождения объемов при цене клиринга
    const supplyVolume = supplyBefore.volume +
      (clearingPrice - supplyBefore.price) *
      (supplyAfter.volume - supplyBefore.volume) /
      (supplyAfter.price - supplyBefore.price);

    const demandVolume = demandBefore.volume +
      (clearingPrice - demandBefore.price) *
      (demandAfter.volume - demandBefore.volume) /
      (demandAfter.price - demandBefore.price);

    // Среднее значение объемов для точки пересечения
    return {
      price: clearingPrice,
      volume: (supplyVolume + demandVolume) / 2
    };
  };

  const CustomTooltip = ({ active, payload, label }) => {
    if (active && payload && payload.length) {
      const interpolateValue = (data, targetPrice) => {
        if (!data || data.length === 0) return null;
        const sorted = [...data].sort((a, b) => a.price - b.price);
        if (targetPrice <= sorted[0].price) return sorted[0].volume;
        if (targetPrice >= sorted[sorted.length - 1].price) return sorted[sorted.length - 1].volume;

        let left, right;
        for (let i = 0; i < sorted.length - 1; i++) {
          if (sorted[i].price <= targetPrice && sorted[i+1].price >= targetPrice) {
            left = sorted[i];
            right = sorted[i+1];
            break;
          }
        }

        const ratio = (targetPrice - left.price) / (right.price - left.price);
        return left.volume + ratio * (right.volume - left.volume);
      };

      const supplyValue = interpolateValue(supply, label);
      const demandValue = interpolateValue(demand, label);

      return (
        <div style={{
          background: 'rgba(30, 30, 47, 0.95)',
          padding: '10px',
          border: '1px solid #a74aff',
          borderRadius: '6px',
          color: 'white',
          fontSize: '14px',
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.3)'
        }}>
          <div style={{ marginBottom: '5px', fontWeight: 'bold' }}>
            {t("main.chart.price")}: {label.toFixed(3)}
          </div>
          {supplyValue !== null && (
            <div style={{ color: '#ff4f81', display: 'flex', alignItems: 'center' }}>
              <div style={{
                width: '10px',
                height: '10px',
                background: '#ff4f81',
                marginRight: '8px',
                borderRadius: '50%'
              }}/>
              {t("main.chart.supply")}: {supplyValue.toFixed(4)}
            </div>
          )}
          {demandValue !== null && (
            <div style={{ color: '#4fc3f7', display: 'flex', alignItems: 'center' }}>
              <div style={{
                width: '10px',
                height: '10px',
                background: '#4fc3f7',
                marginRight: '8px',
                borderRadius: '50%'
              }}/>
              {t("main.chart.demand")}: {demandValue.toFixed(4)}
            </div>
          )}
        </div>
      );
    }
    return null;
  };

  const fetchData = async () => {
    try {
      const clearingPrice = (await getClearingPrice()).price;
      const minPrice = clearingPrice * 0.5;
      const maxPrice = clearingPrice * 1.5;

      const data = await getCurvesData(minPrice, maxPrice);
      setBounds({ min: minPrice, max: maxPrice });

      const parsedSupply = (data.supply || []).map(p => ({
        price: parseFloat(p.price),
        volume: parseFloat(p.volume),
      }));

      const parsedDemand = (data.demand || []).map(p => ({
        price: parseFloat(p.price),
        volume: parseFloat(p.volume),
      }));

      // Используем новую функцию для расчета точки пересечения
      const intersectionPoint = findIntersection(
        parsedSupply,
        parsedDemand,
        parseFloat(clearingPrice)
      );

      setSupply(parsedSupply);
      setDemand(parsedDemand);
      console.log(typeof (intersection.volume));
      setIntersection(intersectionPoint);
    } catch (error) {
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  useInterval(() => {
    fetchData();
  }, 1000);

  if (loading) {
    return <div style={{
      color: 'white',
      textAlign: 'center',
      padding: '20px'
    }}>
      Loading chart data...
    </div>;
  }
  return (
    <div className="chart-section" style={{
      position: 'relative',
      height: '100%',
      width: '100%'
    }}>
      <ResponsiveContainer width="100%" height="100%">
        <LineChart
          margin={{
            top: 10,
            right: 30,
            left: 30,
            bottom: 50,
          }}
        >
          <CartesianGrid strokeDasharray="3 3" stroke="#3a3a5c" opacity={0.5} />

          <XAxis
            allowDataOverflow={true}
            dataKey="price"
            domain={[bounds.min, bounds.max]}
            tick={<CustomizedAxisTick />}
            tickMargin={10}
            tickCount={8}
            stroke="#ffffff"
            tickFormatter={(value) => value.toFixed(3)}
            type="number"
          >
            <Label
              value={t("main.chart.price") + " (USDT)"}
              position="bottom"
              offset={24}
              style={{
                fill: '#ffffff',
                fontSize: '14px',
                fontWeight: 'bold'
              }}
            />
          </XAxis>

          <YAxis
            domain={['auto', 'auto']}
            allowDataOverflow={true}
            type="number"
            tickMargin={15}
            stroke="#ffffff"
            tick={{ fill: '#ffffff', fontSize: 14, fontWeight: 500 }}
            width={80}
          >
            <Label
              value={"BTC " + t("main.chart.volume_hour")}
              angle={-90}
              position="left"
              offset={0}
              style={{
                fill: '#ffffff',
                fontSize: '14px',
                fontWeight: 'bold',
                textAnchor: 'middle'
              }}
            />
          </YAxis>

          <Tooltip
            content={<CustomTooltip />}
            cursor={{ stroke: '#a74aff', strokeWidth: 1, strokeDasharray: '3 3' }}
          />

          <Legend
            verticalAlign="top"
            height={50}
            wrapperStyle={{
              paddingTop: '10px',
              paddingBottom: '10px'
            }}
            iconSize={12}
            iconType="circle"
            formatter={(value) => (
              <span style={{
                color: '#ffffff',
                fontSize: '14px',
                paddingLeft: '5px'
              }}>{value}</span>
            )}
          />

          <Line
            activeDot={false}
            data={supply}
            dataKey="volume"
            stroke="#ff4f81"
            strokeWidth={3}
            name={t("main.chart.supply")}
            dot={false}
            isAnimationActive={false}
          />

          <Line
            data={demand}
            dataKey="volume"
            stroke="#4fc3f7"
            strokeWidth={3}
            name={t("main.chart.demand")}
            dot={false}
            activeDot={false}
            isAnimationActive={false}
          />

            <ReferenceDot
              x={intersection.price}
              y={intersection.volume}
              r={4}
              fill="#00ff00"
              stroke="#fff"
              strokeWidth={2}
            />
               <ReferenceLine
                x={intersection.price}
                stroke="#00ff00"
                strokeWidth={1}
                strokeDasharray="3 3"
                label={{
                  value: `Clearing: ${intersection.price.toFixed(3)}`,
                  position: 'bottom',
                  fill: '#00ff00',
                  fontSize: 12,
                  z_index: 100

                }}
                segment={[
                  { x: intersection.price, y: 0 },
                  { x: intersection.price, y: intersection.volume }
                ]}
              />
              <ReferenceLine
                y={intersection.volume}
                stroke="#00ff00"
                strokeWidth={1}
                strokeDasharray="3 3"
                label={{
                  value: `Volume: ${intersection.volume.toFixed(4)}`,
                  position: 'left',
                  fill: '#00ff00',
                  fontSize: 12,
                  z_index: 100
                }}
                segment={[
                  { x: 0, y: intersection.volume },
                  { x: intersection.price, y: intersection.volume }
                ]}
              />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}