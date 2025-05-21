"use client";
import React, { useState } from "react";
import {
  LaptopOutlined,
  NotificationOutlined,
  UserOutlined,
  DashboardOutlined,
  UsergroupAddOutlined,
  GlobalOutlined,
  DollarOutlined,
} from "@ant-design/icons";
import { Breadcrumb, Dropdown, Layout, Menu, theme } from "antd";
import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";
import { Avatar, Space } from "antd";
import { logOut, useUser } from "@/lib/session";
import { Typography } from "antd";

const { Title } = Typography;

function UserMenu({ children }) {
  const items = [
    {
      label: "退出登录",
      key: "logout",
    },
  ];

  const onClick = ({ key }) => {
    console.log(key);
    if (key == "logout") {
      logOut();
    }
  };

  return (
    <Dropdown menu={{ items, onClick }} trigger={"click"}>
      {children}
    </Dropdown>
  );
}

const { Header, Content, Sider } = Layout;

const items2 = [
  {
    key: "dashboard",
    label: "仪表盘",
    icon: <DashboardOutlined />,
  },
  {
    key: "promotion",
    label: "推广明细",
    icon: <DollarOutlined />,
  },
  {
    key: "domains",
    label: "域名管理",
    icon: <GlobalOutlined />,
  },
  {
    key: "team",
    label: "团队信息",
    icon: <UsergroupAddOutlined />,
  },
  {
    key: "me",
    label: "个人中心",
    icon: <UserOutlined />,
  },
];

export default function DashboardLayout({ children }) {
  const router = useRouter();
  const path = usePathname(); // 获取当前路由路径，如 "/login"
  const { user, error, isLoading } = useUser();
  const [selectedKeys, setSelectedKeys] = useState([]);

  function handleSelected({ item, key, keyPath, selectedKeys, domEvent }) {
    console.log({ item, key, keyPath, selectedKeys });
    setSelectedKeys(selectedKeys);
    router.push("/" + selectedKeys.join("/"));
  }

  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  if (isLoading) {
    return;
  }
  if (error) {
    console.log(error);
    router.push("/login"); // temp
    return;
  }

  return (
    <Layout style={{ height: "100vh" }}>
      <Header style={{ display: "flex", alignItems: "center" }}>
        <div className="demo-logo" />
        {/* <Title
          level={3}
          style={{
            margin: 0,
            color: "white",
            lineHeight: "inherit", // 继承 Header 的高度
            paddingLeft: 16,
          }}
        >
          代理系统后台
        </Title> */}
        <UserMenu>
          {user ? (
            <Avatar style={{ marginLeft: "auto", backgroundColor: "#87d068" }}>
              {user.username}
            </Avatar>
          ) : (
            <Avatar icon={<UserOutlined />} style={{ marginLeft: "auto" }} />
          )}
        </UserMenu>
      </Header>
      <Layout>
        <Sider
          width={200}
          style={{ background: colorBgContainer }}
          breakpoint="lg"
          collapsedWidth="0"
        >
          <Menu
            mode="inline"
            defaultSelectedKeys={["dashboard"]}
            defaultOpenKeys={["dashboard"]}
            style={{ height: "100%", borderRight: 0 }}
            items={items2}
            onSelect={handleSelected}
            selectedKeys={selectedKeys}
          />
        </Sider>
        <Layout style={{ padding: "0 24px 24px" }}>
          <Breadcrumb
            style={{ margin: "18px 0" }}
            items={[
              { title: "主页" },
              ...selectedKeys.map((key) => ({
                title: items2.find((e) => e.key == key).label,
              })),
            ]}
          />
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
