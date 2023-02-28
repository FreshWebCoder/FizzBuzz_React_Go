import React, { useEffect, useMemo, useState } from 'react';
import { LoadingOutlined } from '@ant-design/icons';
import { Button, Space, Typography, notification } from 'antd';
import axios from 'axios';
import { debounce } from 'lodash';

import './App.css';

const { Text, Title } = Typography;
const { REACT_APP_API_SERVER: API_SERVER } = process.env;

const App: React.FC = () => {
  const [count, setCount] = useState(1);
  const [message, setMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const debounceCall = useMemo(
    () => debounce((val: number) => {
      setIsSubmitting(true);
      setMessage("");

      axios
        .post(`${API_SERVER}/fizzbuzz`, { count: val })
        .then(({ data }) => {
          setIsSubmitting(false);
          setMessage(data.message)
        })
        .catch((err) => {
          setIsSubmitting(false);
          console.error("API Error: ", err);
          notification.error({
            message: "Error",
            description: "Something went wrong!"
          })
        })
      },
      500
    ), []);

  useEffect(() => {
    debounceCall(count);

    return debounceCall.cancel;
  }, [count, debounceCall]);

  const increaseCount = () => {
    setCount(count + 1);
  }

  return (
    <Space className="App" direction="vertical" size={54} align="center">
      <Text className="text-count">Your count<br />{count}</Text>

      <Button type="primary" className="btn-increase" onClick={increaseCount}>Push me!</Button>

      {message ? (
        <Title level={1} className="text-message">{message}</Title>
      ) : isSubmitting && (
        <LoadingOutlined style={{ fontSize: 36 }} />
      )}
    </Space>
  );
}

export default App;
