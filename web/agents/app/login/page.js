"use client";
import React, { useState } from "react";
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Checkbox, Form, Input, Flex } from "antd";
import { Card } from "antd";
import { Typography } from "antd";
import { Alert } from "antd";

function LoginForm() {
  const [err, setErr] = useState(null);

  const onFinish = async (values) => {
    const res = await fetch(
      process.env.NEXT_PUBLIC_API_BASE + "/api/auth/login",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: values.username,
          password: values.password,
        }),
      }
    );
    const json = await res.json();

    if (!res.ok) {
      setErr(JSON.stringify(json));
      return;
    }

    localStorage.setItem("token", json.token);
    window.location.href = "/dashboard";
  };

  return (
    <>
      {err && (
        <Alert
          message={err}
          type="error"
          showIcon
          style={{ marginBottom: "10px" }}
        />
      )}

      <Form
        name="login"
        initialValues={{ remember: true }}
        style={{ maxWidth: 360 }}
        onFinish={onFinish}
      >
        <Form.Item
          name="username"
          rules={[{ required: true, message: "Please input your Username!" }]}
        >
          <Input prefix={<UserOutlined />} placeholder="用户名" />
        </Form.Item>

        <Form.Item
          name="password"
          rules={[{ required: true, message: "Please input your Password!" }]}
        >
          <Input prefix={<LockOutlined />} type="password" placeholder="密码" />
        </Form.Item>

        <Form.Item>
          <Flex justify="space-between" align="center">
            <Form.Item name="remember" valuePropName="checked" noStyle>
              <Checkbox>记住我</Checkbox>
            </Form.Item>
            <a href="">忘记密码</a>
          </Flex>
        </Form.Item>

        <Form.Item>
          <Button block type="primary" htmlType="submit">
            登录
          </Button>
        </Form.Item>
      </Form>
    </>
  );
}
const App = () => {
  return (
    <Flex align="center" justify="center" style={{ height: "100vh" }}>
      <Card style={{ width: 360 }}>
        <Typography.Title level={3} style={{ textAlign: "center" }}>
          您好！请登录
        </Typography.Title>
        <LoginForm></LoginForm>
      </Card>
    </Flex>
  );
};
export default App;
