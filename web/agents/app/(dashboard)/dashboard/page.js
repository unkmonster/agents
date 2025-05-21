"use client";
import { useTotalCommission, useUser } from "@/lib/session";
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
} from "antd";
import Meta from "antd/es/card/Meta";
import Title from "antd/es/typography/Title";

function StatCard({ title, value, prefix }) {
  return (
    <Card variant="">
      <Statistic value={value} title={title} prefix={prefix} />
    </Card>
  );
}

function CommissionCard() {
  const { commission, isLoading, error } = useTotalCommission();
  if (isLoading) {
    return <Spin />;
  }
  if (error) {
    return error;
  }

  return (
    <Card title="充值数据">
      <Flex gap={"middle"} vertical>
        <Statistic
          title="今日充值"
          value={commission.todayCommission / 100}
          prefix={"￥"}
        />
        <Statistic
          title="全部充值"
          value={commission.totalCommission / 100}
          prefix={"￥"}
        />
        <Statistic
          title="已提现金额"
          value={commission.settledCommission / 100}
          prefix={"￥"}
        />
        <Statistic
          title="账户余额"
          value={
            (commission.totalCommission - commission.settledCommission) / 100
          }
          prefix={"￥"}
        />
      </Flex>
    </Card>
  );
}

export default function Main() {
  const { user, isLoading, error } = useUser();
  const commRes = useTotalCommission();

  if (isLoading || commRes.isLoading) {
    return <Spin />;
  }
  if (error || commRes.error) {
    return (error || commRes.error).toString();
  }

  return (
    <>
      <Row gutter={[32, 32]}>
        <Col xs={24} sm={12} md={8} xl={6}>
          <StatCard
            title={"今日注册"}
            value={commRes.commission.todayRegistrationCount}
          />
        </Col>

        <Col xs={24} sm={12} md={8} xl={6}>
          <StatCard
            title={"今日充值"}
            value={commRes.commission.todayCommission}
            prefix={"￥"}
          />
        </Col>

        <Col xs={24} sm={12} md={8} xl={6}>
          <StatCard
            title={"累计充值"}
            value={commRes.commission.totalCommission}
            prefix={"￥"}
          />
        </Col>

        <Col xs={24} sm={12} md={8} xl={6}>
          <StatCard
            title={"账户余额"}
            value={
              commRes.commission.totalCommission -
              commRes.commission.settledCommission
            }
            prefix={"￥"}
          />
        </Col>
      </Row>
    </>
  );
}
