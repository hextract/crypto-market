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
  ReferenceDot,
} from 'recharts';
import { getCurvesData } from '../api/matchingEngineService';

export default function MarketChart() {
  const [supply, setSupply] = useState([]);
  const [demand, setDemand] = useState([]);
  const [clearingPrice, setClearingPrice] = useState(null);
  const [intersection, setIntersection] = useState(null);

  const leftBoundary = 400;
  const rightBoundary = 600;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getCurvesData(leftBoundary, rightBoundary);
        setSupply(data.supply || []);
        setDemand(data.demand || []);
        setClearingPrice(data.clearingPrice || 500); // fallback if missing

        // Обработка и фильтрация
        const parsedSupply = (data.supply || []).map(p => ({
          price: parseFloat(p.price),
          volume: parseFloat(p.volume),
        })).filter(p => p.price >= leftBoundary && p.price <= rightBoundary);

        const parsedDemand = (data.demand || []).map(p => ({
          price: parseFloat(p.price),
          volume: parseFloat(p.volume),
        })).filter(p => p.price >= leftBoundary && p.price <= rightBoundary);

        const price = data.clearingPrice || 500;

        const supplyPoint = parsedSupply.find(p => p.price >= price) || parsedSupply[parsedSupply.length - 1];
        const demandPoint = parsedDemand.find(p => p.price >= price) || parsedDemand[parsedDemand.length - 1];
        const avgVolume = (supplyPoint.volume + demandPoint.volume) / 2;

        setIntersection({ price, volume: avgVolume });
      } catch (error) {
        console.error('Error fetching curves data:', error);
      }
    };

    fetchData();
  }, []);

  const parsedSupply = supply.map(p => ({
    price: parseFloat(p.price),
    volume: parseFloat(p.volume),
  })).filter(p => p.price >= leftBoundary && p.price <= rightBoundary);

  const parsedDemand = demand.map(p => ({
    price: parseFloat(p.price),
    volume: parseFloat(p.volume),
  })).filter(p => p.price >= leftBoundary && p.price <= rightBoundary);

  if (!intersection || clearingPrice === null) {
    return <div>Loading chart...</div>;
  }

  return (
    <div className="chart-section">
      <ResponsiveContainer width="95%" height="90%">
        <LineChart>
          <CartesianGrid strokeDasharray="3 3" stroke="#8884d8" />
          <XAxis dataKey="price" type="number" domain={[leftBoundary, rightBoundary]} stroke="#ffffff" />
          <YAxis type="number" stroke="#ffffff" />
          <Tooltip />
          <Legend />
          <Line
            data={parsedSupply}
            type="monotone"
            dataKey="volume"
            stroke="#ff4f81"
            name="Supply"
            dot={false}
          />
          <Line
            data={parsedDemand}
            type="monotone"
            dataKey="volume"
            stroke="#4fc3f7"
            name="Demand"
            dot={false}
          />
          <ReferenceDot
            x={intersection.price}
            y={intersection.volume}
            r={6}
            fill="#00ff00"
            stroke="#fff"
            label={{ value: '', position: 'top', fill: '#fff' }}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
