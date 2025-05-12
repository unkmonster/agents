"use client";
import React from "react";
import {
  LaptopOutlined,
  NotificationOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { Breadcrumb, Layout, Menu, theme } from "antd";
import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";

const { Header, Content, Sider } = Layout;
const items1 = [
  {
    key: 1,
    label: "主页",
  },
];
const items2 = [
  {
    key: "/dashboard",
    label: "仪表盘",
  },
  {
    key: "/promotion",
    label: "推广明细",
  },
  {
    key: "/domains",
    label: "域名管理",
  },
  {
    key: "/team",
    label: "团队信息",
  },
];

export default function DashboardLayout({ children }) {
  const router = useRouter();
  const path = usePathname(); // 获取当前路由路径，如 "/login"

  function handleSelected({ item, key, keyPath, selectedKeys, domEvent }) {
    console.log({ item, key, keyPath, selectedKeys, domEvent });
    router.push(key);
  }

  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  return (
    <Layout style={{ height: "100vh" }}>
      <Header style={{ display: "flex", alignItems: "center" }}>
        <div className="demo-logo" />
        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={["1"]}
          items={items1}
          style={{ flex: 1, minWidth: 0 }}
        />
      </Header>
      <Layout>
        <Sider width={200} style={{ background: colorBgContainer }}>
          <Menu
            mode="inline"
            defaultSelectedKeys={["/dashboard"]}
            defaultOpenKeys={["/dashboard"]}
            style={{ height: "100%", borderRight: 0 }}
            items={items2}
            onSelect={handleSelected}
            selectedKeys={[path]}
          />
        </Sider>
        <Layout style={{ padding: "0 24px 24px" }}>
          <Content
            style={{
              padding: 24,
              margin: 0,
              minHeight: 280,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            {children}
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
}
