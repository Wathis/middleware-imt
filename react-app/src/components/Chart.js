import React from 'react';
import { LineChart, Line, XAxis, YAxis, Label, ResponsiveContainer } from 'recharts';
import Title from './Title';

export default function Chart(props) {
  return (
    <React.Fragment>
      <Title>{props.title}{props.airportId}</Title>
      <ResponsiveContainer>
        <LineChart
          data={formatData(props.data, props.airportId, props.measure)}
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
                {props.xAxis}
            </Label>
          </YAxis>
          <Line type="monotone" dataKey="amount" stroke="#556CD6" dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
}

function formatData(data, filterAirportId, filterMeasure) {
    
    // FILTER DATA
    let measures = data.filter(function(elem) {
        return elem.airportId === filterAirportId && elem.measureType === filterMeasure
    })

    // FORMAT FOR GRAPH
    let res = measures.map(function(elem){
        return createData(new Date(elem.timestamp*1000).toLocaleString(), elem.measureValue)
    })
    // console.log(measures)
    return res
}

function createData(time, amount) {
    return { time, amount };
  }
  