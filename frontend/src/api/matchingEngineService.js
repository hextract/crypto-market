import axios from 'axios';

const API_BASE = process.env.REACT_APP_MATCHING_ENGINE_API_BASE_URL;
const API_TIMEOUT = process.env.REACT_APP_API_TIMEOUT;

const api = axios.create({
  baseURL: API_BASE,
  timeout: parseInt(API_TIMEOUT),
});

// Добавляем интерцептор для авторизации
api.interceptors.request.use(config => {
  const token = getAuthToken();
  if (token) {
    config.headers['api_key'] = token;
  }
  return config;
});


export const getCurvesData = async (left_boundary_price, right_boundary_price) => {
  try {
    // const response = await api.get('/curves', {
    //   left_boundary_price,
    //   right_boundary_price,
    // });
    const response1 = {
      data: {
        supply: [{'price': '0', 'volume': '23244'}, {'price': '200.576', 'volume': '23244'}, {'price': '312.728', 'volume': '23209.8'}, {'price': '320.615', 'volume': '23175.3'}, {'price': '329.278', 'volume': '22984.3'}, {'price': '347.878', 'volume': '22502.1'}, {'price': '360.837', 'volume': '22142'}, {'price': '362.256', 'volume': '22127.7'}, {'price': '373.171', 'volume': '21961.8'}, {'price': '381.301', 'volume': '21786.7'}, {'price': '382.991', 'volume': '21749.9'}, {'price': '384.366', 'volume': '21706.9'}, {'price': '386.211', 'volume': '21645.2'}, {'price': '387.013', 'volume': '21618.4'}, {'price': '394.854', 'volume': '21331.6'}, {'price': '406.045', 'volume': '20889.1'}, {'price': '407.683', 'volume': '20812.2'}, {'price': '408.356', 'volume': '20778.4'}, {'price': '429.679', 'volume': '19465.3'}, {'price': '436.881', 'volume': '18993.1'}, {'price': '448.164', 'volume': '17960.3'}, {'price': '451.531', 'volume': '17646.9'}, {'price': '452.399', 'volume': '17570.5'}, {'price': '455.062', 'volume': '17366.6'}, {'price': '460.449', 'volume': '16937.7'}, {'price': '464.114', 'volume': '16394.6'}, {'price': '466.3', 'volume': '16084.5'}, {'price': '466.942', 'volume': '15989.4'}, {'price': '468.252', 'volume': '15749.3'}, {'price': '469.807', 'volume': '15570.9'}, {'price': '470.454', 'volume': '15496'}, {'price': '472.871', 'volume': '15158'}, {'price': '481.127', 'volume': '14218.1'}, {'price': '481.178', 'volume': '14211.2'}, {'price': '482.398', 'volume': '14059'}, {'price': '483.549', 'volume': '13955.7'}, {'price': '485.495', 'volume': '13781.7'}, {'price': '490.418', 'volume': '13324.2'}, {'price': '495.93', 'volume': '12797.7'}, {'price': '504.928', 'volume': '11906.9'}, {'price': '510.34', 'volume': '11344.4'}, {'price': '511.208', 'volume': '11257.2'}, {'price': '511.861', 'volume': '11207.5'}, {'price': '512.999', 'volume': '11118.8'}, {'price': '513.571', 'volume': '11071.4'}, {'price': '515.158', 'volume': '10946.3'}, {'price': '516.253', 'volume': '10857.4'}, {'price': '516.872', 'volume': '10803.5'}, {'price': '521.808', 'volume': '10367.4'}, {'price': '525.079', 'volume': '10102.7'}, {'price': '526.213', 'volume': '10015.6'}, {'price': '527.249', 'volume': '9957.11'}, {'price': '531.345', 'volume': '8737.01'}, {'price': '538.36', 'volume': '8341.13'}, {'price': '543.481', 'volume': '8067.39'}, {'price': '556.414', 'volume': '7308.11'}, {'price': '560.923', 'volume': '7029.82'}, {'price': '561.259', 'volume': '7009.58'}, {'price': '564.743', 'volume': '6672.54'}, {'price': '569.901', 'volume': '6068.22'}, {'price': '575.025', 'volume': '5484.32'}, {'price': '575.94', 'volume': '5366.65'}, {'price': '581.713', 'volume': '4641.04'}, {'price': '582.142', 'volume': '4602.75'}, {'price': '587.521', 'volume': '4201.76'}, {'price': '588.771', 'volume': '4110.5'}, {'price': '588.858', 'volume': '4104.5'}, {'price': '598.753', 'volume': '3623.38'}, {'price': '599.683', 'volume': '3584.07'}, {'price': '602.259', 'volume': '3481.19'}, {'price': '610.971', 'volume': '3027.71'}, {'price': '614.514', 'volume': '2860.36'}, {'price': '631.872', 'volume': '2005.79'}, {'price': '638.854', 'volume': '1675.01'}, {'price': '643.092', 'volume': '1495.18'}, {'price': '646.704', 'volume': '1354.46'}, {'price': '660.907', 'volume': '813.104'}, {'price': '664.766', 'volume': '688.326'}, {'price': '664.814', 'volume': '686.916'}, {'price': '666.218', 'volume': '653.22'}, {'price': '670.261', 'volume': '566.675'}, {'price': '685.117', 'volume': '428.577'}, {'price': '690.227', 'volume': '382.312'}, {'price': '708.519', 'volume': '272.845'}, {'price': '754.317', 'volume': '99.7851'}, {'price': '777.477', 'volume': '13.4675'}, {'price': '785.264', 'volume': '0'}, {'price': '979.073', 'volume': '0'}],
        demand: [{'price': '0', 'volume': '0'}, {'price': '242.272', 'volume': '0'}, {'price': '261.176', 'volume': '98.2506'}, {'price': '264.147', 'volume': '114.865'}, {'price': '284.139', 'volume': '245.309'}, {'price': '286.663', 'volume': '276.762'}, {'price': '313.831', 'volume': '474.165'}, {'price': '315.683', 'volume': '487.892'}, {'price': '319.683', 'volume': '542.154'}, {'price': '325.3', 'volume': '636.365'}, {'price': '330.299', 'volume': '740.725'}, {'price': '356.17', 'volume': '1407.63'}, {'price': '357.167', 'volume': '1440.2'}, {'price': '360.97', 'volume': '1540.94'}, {'price': '366.486', 'volume': '1693.74'}, {'price': '371.111', 'volume': '1827.24'}, {'price': '375.241', 'volume': '1948.11'}, {'price': '375.476', 'volume': '1954.77'}, {'price': '381.956', 'volume': '2143.55'}, {'price': '384.185', 'volume': '2418.24'}, {'price': '388.999', 'volume': '3039.13'}, {'price': '392.117', 'volume': '3147.96'}, {'price': '393.494', 'volume': '3208.59'}, {'price': '398.798', 'volume': '3905.59'}, {'price': '400.513', 'volume': '4133.99'}, {'price': '400.816', 'volume': '4175.62'}, {'price': '406.532', 'volume': '4462.91'}, {'price': '408.699', 'volume': '4576.06'}, {'price': '409.253', 'volume': '4606.3'}, {'price': '413.26', 'volume': '4841.3'}, {'price': '418.306', 'volume': '5137.42'}, {'price': '423.705', 'volume': '5470.9'}, {'price': '432.309', 'volume': '5943.3'}, {'price': '433.389', 'volume': '6010.55'}, {'price': '433.88', 'volume': '6041.66'}, {'price': '435.769', 'volume': '6170.2'}, {'price': '437.1', 'volume': '6264.19'}, {'price': '438.698', 'volume': '6367.55'}, {'price': '440.404', 'volume': '6486.79'}, {'price': '442.153', 'volume': '6616.43'}, {'price': '443.985', 'volume': '6735.53'}, {'price': '452.863', 'volume': '7315.32'}, {'price': '453.407', 'volume': '7350.9'}, {'price': '458.384', 'volume': '7746.69'}, {'price': '463.991', 'volume': '8191.73'}, {'price': '464.834', 'volume': '8255.17'}, {'price': '467.224', 'volume': '8424.29'}, {'price': '469.349', 'volume': '8575.76'}, {'price': '471.62', 'volume': '8726.42'}, {'price': '476.919', 'volume': '9220.69'}, {'price': '479.837', 'volume': '9509.27'}, {'price': '482.12', 'volume': '9747.32'}, {'price': '482.779', 'volume': '9820'}, {'price': '490.658', 'volume': '10761.3'}, {'price': '491.453', 'volume': '10863.1'}, {'price': '491.825', 'volume': '10913'}, {'price': '491.886', 'volume': '10924.1'}, {'price': '496.413', 'volume': '11766.8'}, {'price': '496.752', 'volume': '11830.7'}, {'price': '496.927', 'volume': '11863.5'}, {'price': '497.702', 'volume': '12004.2'}, {'price': '500.241', 'volume': '12490.2'}, {'price': '505.128', 'volume': '13194.9'}, {'price': '507.302', 'volume': '13522.1'}, {'price': '510.453', 'volume': '13911.5'}, {'price': '513.189', 'volume': '14263'}, {'price': '519.737', 'volume': '15583.8'}, {'price': '520.92', 'volume': '15824.9'}, {'price': '521.533', 'volume': '15941.1'}, {'price': '521.647', 'volume': '15964.1'}, {'price': '523.06', 'volume': '16246.2'}, {'price': '523.533', 'volume': '16341.6'}, {'price': '523.691', 'volume': '16373'}, {'price': '523.987', 'volume': '16429.1'}, {'price': '528.605', 'volume': '16966.3'}, {'price': '530.603', 'volume': '17211.8'}, {'price': '533.503', 'volume': '17553.1'}, {'price': '537.644', 'volume': '18059.5'}, {'price': '538.611', 'volume': '18177.3'}, {'price': '545.512', 'volume': '19001.6'}, {'price': '545.716', 'volume': '19024.9'}, {'price': '558.931', 'volume': '20445.2'}, {'price': '560.477', 'volume': '20608.3'}, {'price': '562.729', 'volume': '20832.8'}, {'price': '575.934', 'volume': '22168.9'}, {'price': '576.389', 'volume': '22216.9'}, {'price': '579.415', 'volume': '22535'}, {'price': '580.14', 'volume': '22607.1'}, {'price': '580.444', 'volume': '22636.8'}, {'price': '583.746', 'volume': '22953.5'}, {'price': '588.721', 'volume': '23384.7'}, {'price': '614.514', 'volume': '23400.36'}],
        clearing_price: 500,
      }
    };
    return response1.data;
  } catch (error) {
    throw error;
  }
};

// Вспомогательная функция для получения токена
const getAuthToken = () => {
  const cookies = document.cookie.split(';');
  const tokenCookie = cookies.find(c => c.trim().startsWith('token='));
  return tokenCookie ? tokenCookie.split('=')[1] : null;
};