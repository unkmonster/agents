"use client";
import React, { useEffect, useState } from "react";
import {
  LaptopOutlined,
  NotificationOutlined,
  UserOutlined,
  DashboardOutlined,
  UsergroupAddOutlined,
  GlobalOutlined,
  DollarOutlined,
} from "@ant-design/icons";
import { Breadcrumb, Dropdown, Layout, Menu, Result, Spin, theme } from "antd";
import { useRouter } from "next/navigation";
import { usePathname } from "next/navigation";
import { Avatar, Space } from "antd";
import { logOut, useUser, useUserId } from "@/lib/session";
import { Typography } from "antd";
import { UserContext } from "@/lib/user_context";

const { Title } = Typography;

function UserMenu({ children }) {
  const router = useRouter();
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
      router.push("/login");
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
  const { user, error, isLoading } = useUser();
  const path = usePathname();
  const keys = path.split("/").filter((v) => v);
  function handleSelected({ item, key, keyPath, selectedKeys, domEvent }) {
    //console.log({ item, key, keyPath, selectedKeys });
    //setSelectedKeys(selectedKeys);
    router.push("/" + selectedKeys.join("/"));
  }
  const [err, setErr] = useState();
  const [jumping, setJumping] = useState(false);

  useEffect(() => {
    if (!error) {
      return;
    }
    setErr(error);
    if (error.isUnauthorized()) {
      setJumping(true);
      setTimeout(() => router.push("/login"), 1000);
    }
  }, [error]);

  // getUser
  if (isLoading) {
    return <Spin fullscreen tip="正在加载用户信息" />;
  }

  if (jumping) {
    return <Spin fullscreen tip="未登录/登陆失效 正在跳转至登陆界面" />;
  }
  // getUser
  if (err) {
    return (
      <Result
        status={"error"}
        title={"获取用户信息失败"}
        subTitle={JSON.stringify(err)}
      />
    );
  }

  function itemRender(currentRoute, params, items, paths) {
    const isLast = currentRoute?.path === items[items.length - 1]?.path;

    return isLast ? (
      <span>{currentRoute.title}</span>
    ) : (
      <a href={`/${paths.join("/")}`}>{currentRoute.title}</a>
    );
  }

  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Header style={{ display: "flex", alignItems: "center", height: "48px" }}>
        <div className="demo-logo">
          <Space direction="vertical" align="center">
            <span
              style={{
                color: "white",
                fontSize: "14px",
                fontWeight: "500",
                userSelect: "none",
              }}
            >
              AGENTS
            </span>
          </Space>
        </div>

        <UserMenu>
          {user ? (
            <Avatar
              style={{
                marginLeft: "auto",
                backgroundColor: "#87d068",
                userSelect: "none",
              }}
            >
              {user.user?.username[0].toUpperCase()}
            </Avatar>
          ) : (
            <Avatar
              icon={<UserOutlined />}
              style={{ marginLeft: "auto", userSelect: "none" }}
            />
          )}
        </UserMenu>
      </Header>
      <Layout>
        <Sider width={200} breakpoint="lg" collapsedWidth="0">
          <Menu
            mode="inline"
            defaultSelectedKeys={["dashboard"]}
            defaultOpenKeys={["dashboard"]}
            style={{ height: "100%", borderRight: 0 }}
            items={items2}
            onSelect={handleSelected}
            selectedKeys={keys}
          />
        </Sider>
        <Layout style={{ padding: "0 24px 24px" }}>
          <Breadcrumb
            style={{ margin: "18px 0" }}
            itemRender={itemRender}
            items={[
              { title: "主页" },
              ...keys.map((key) => ({
                title: items2.find((e) => e.key == key)?.label,
              })),
            ]}
          />
          <Content
            style={{
              //padding: 24,
              margin: 0,
              //minHeight: 280,
              // background: colorBgContainer,
              //borderRadius: borderRadiusLG,
            }}
          >
            <UserContext.Provider value={user?.user}>
              {children}
            </UserContext.Provider>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
}
