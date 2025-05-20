"use client";
import { Card, Button, Col, Row, Statistic } from "antd";

export default function Main() {
  return (
    <>
      <Row gutter={32}>
        <Col xs={24} sm={12} md={8} xl={6}>
          <Card>
            <Statistic title="今日充值" value={30} prefix="￥"></Statistic>
          </Card>
        </Col>
        <Col xs={24} sm={12} md={8} xl={6}>
          <Card>
            <Statistic title="全部充值" value={30} prefix="￥"></Statistic>
          </Card>
        </Col>
        <Col xs={24} sm={12} md={8} xl={6}>
          <Card>
            <Statistic title="已提现金额" value={30} prefix="￥"></Statistic>
          </Card>
        </Col>
        <Col xs={24} sm={12} md={8} xl={6}>
          <Card>
            <Statistic title="账户余额" value={30} prefix="￥"></Statistic>
          </Card>
        </Col>
      </Row>
    </>
  );
}
