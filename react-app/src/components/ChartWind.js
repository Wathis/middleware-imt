import React from 'react';
import { LineChart, Line, XAxis, YAxis, Label, ResponsiveContainer } from 'recharts';
import Title from './Title';

// Generate Sales Data
function createData(time, amount) {
  return { time, amount };
}

const data = [
  createData('2019-10-16 00:00', 10),
  createData('2019-10-16 03:00', 10.8),
  createData('2019-10-16 06:00', 10.4),
  createData('2019-10-16 09:00', 12.1),
  createData('2019-10-16 12:00', 18.3),
  createData('2019-10-16 15:00', 21.9),
  createData('2019-10-16 18:00', 21.4),
  createData('2019-10-16 21:00', 15),
];

export default function Chart() {
  return (
    <React.Fragment>
      <Title>Evolution du vent</Title>
      <ResponsiveContainer>
        <LineChart
          data={data}
          margin={{
            top: 16,
            right: 16,
            bottom: 0,
            left: 24,
          }}
        >
          <XAxis dataKey="time" />
          <YAxis>
            <Label angle={270} position="left" style={{ textAnchor: 'middle' }}>
              Vent (km/h)
            </Label>
          </YAxis>
          <Line type="monotone" dataKey="amount" stroke="#556CD6" dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
}