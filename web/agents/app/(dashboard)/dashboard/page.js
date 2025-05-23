"use client";
import {
  myfetch,
  useTodayCommission,
  useTotalCommission,
  useUser,
  useUserId,
} from "@/lib/session";
import {
  Card,
  Button,
  Col,
  Row,
  Statistic,
  Avatar,
  Typography,
  Flex,
  Spin,
  Space,
  Dropdown,
  Empty,
} from "antd";
import { DownOutlined, SmileOutlined } from "@ant-design/icons";

import Meta from "antd/es/card/Meta";
import Title from "antd/es/typography/Title";
import { Column, Line } from "@ant-design/plots";
import useSWR from "swr";
import { useEffect, useState } from "react";

function SevenDaysRechargeChart() {
  const userId = useUserId();
  const [dataSource, setDataSource] = useState([]);
  const [days, setDays] = useState(7);
  const { data, error, isLoading } = useSWR(
    `/v1/users/${userId}/commissions?limit=${days}`,
    myfetch
  );

  useEffect(() => {
    if (!data) {
      return;
    }

    const rechargeDs = data.commissions.map((com) => ({
      date: new Date(com.date),
      amount:
        (parseInt(com.indirectRechargeAmount) +
          parseInt(com.directRechargeAmount)) /
        100,
      type: "充值量",
    }));

    const registerDs = data.commissions.map((com) => ({
      date: new Date(com.date),
      amount:
        parseInt(com.indirectRegistrationCount) +
        parseInt(com.directRegistrationCount),
      type: "注册量",
    }));

    setDataSource([...registerDs, ...rechargeDs]);
  }, [data]);

  if (error) {
    return JSON.stringify(error);
  }

  const config = {
    data: dataSource,
    xField: "date",
    yField: "amount",
    point: {
      shapeField: "square",
      sizeField: 2,
    },
    interaction: {
      tooltip: {
        marker: false,
      },
    },
    style: {
      lineWidth: 2,
    },
    smooth: true,
    colorField: "type",
    //title: "推广趋势",
    itemLabelText: "你好",
  };

  const items = [
    {
      key: "7",
      label: "7 日趋势",
    },
    {
      key: "30",
      label: "30 日趋势",
    },
  ];

  return (
    <Card
      title="推广趋势"
      loading={isLoading}
      extra={
        <Dropdown
          menu={{
            items,
            selectable: true,
            defaultSelectedKeys: [String(days)],
            onClick: ({ _, key }) => setDays(key),
          }}
          trigger={["click"]}
        >
          <Typography.Link>
            <Space>
              {`${days} 日趋势`} <DownOutlined />
            </Space>
          </Typography.Link>
        </Dropdown>
      }
    >
      {!dataSource ? <Empty /> : <Line {...config} />}
    </Card>
  );
}

function StatCard({ title, value, prefix, actions }) {
  return (
    <Card variant="" actions={actions}>
      <Statistic value={value} title={title} prefix={prefix} />
    </Card>
  );
}

export default function Main() {
  const commRes = useTotalCommission();
  const todayCommRes = useTodayCommission();

  return (
    <Space direction="vertical" size={"middle"} style={{ width: "100%" }}>
      <Title level={3}>总览</Title>
      <Row
        gutter={[
          { xs: 8, sm: 16, md: 24, lg: 32 },
          { xs: 8, sm: 16, md: 24, lg: 32 },
        ]}
      >
        <Col xs={24} sm={12} md={8} xl={6}>
          <Spin spinning={todayCommRes.isLoading}>
            <StatCard
              title={"今日注册"}
              value={
                todayCommRes.data &&
                parseInt(
                  todayCommRes.data?.commissions[0]?.directRegistrationCount ||
                    0
                ) +
                  parseInt(
                    todayCommRes.data?.commissions[0]
                      ?.indirectRegistrationCount || 0
                  )
              }
              loading={todayCommRes.isLoading}
            />
          </Spin>
        </Col>

        <Col xs={24} sm={12} md={8} xl={6}>
          <Spin spinning={todayCommRes.isLoading}>
            <StatCard
              title={"今日充值"}
              value={
                todayCommRes.data &&
                (parseInt(
                  todayCommRes.data.commissions[0].directRechargeAmount
                ) +
                  parseInt(
                    todayCommRes.data.commissions[0].indirectRechargeAmount
                  )) /
                  100
              }
              prefix={"￥"}
              loading={todayCommRes.isLoading}
            />
          </Spin>
        </Col>

        <Col xs={24} sm={12} md={8} xl={6}>
          <Spin spinning={commRes.isLoading}>
            <StatCard
              title={"累计充值"}
              value={
                commRes.commission && commRes.commission.totalCommission / 100
              }
              prefix={"￥"}
              loading={commRes.isLoading}
            />
          </Spin>
        </Col>

        <Col xs={24} sm={12} md={8} xl={6}>
          <Spin spinning={commRes.isLoading}>
            <StatCard
              title={"账户余额"}
              value={
                commRes.commission &&
                (commRes.commission.totalCommission -
                  commRes.commission.settledCommission) /
                  100
              }
              prefix={"￥"}
              loading={commRes.isLoading}
              actions={[<Button>提现</Button>]}
            />
          </Spin>
        </Col>
      </Row>

      <Row
        gutter={[
          { xs: 8, sm: 16, md: 24, lg: 32 },
          { xs: 8, sm: 16, md: 24, lg: 32 },
        ]}
      >
        <Col xs={24} sm={12}>
          <SevenDaysRechargeChart />
        </Col>
      </Row>
    </Space>
  );
}
