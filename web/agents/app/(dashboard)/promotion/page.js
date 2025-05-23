"use client";
import { myfetch, useUser, useUserId } from "@/lib/session";
import { Card, Spin, Table } from "antd";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import useSWR from "swr";

export default function Main() {
  const router = useRouter();
  const userId = useUserId();
  const [dataSource, setDataSource] = useState([]);
  const userRes = useUser();
  const isDirectAgent = userRes.user && userRes.user.level == 2;
  const commListRes = useSWR(`/v1/users/${userId}/commissions`, myfetch);

  useEffect(() => {
    if (!userId) {
      //router.push("/login");
      return;
    }

    if (!commListRes.data) {
      return;
    }

    const ds = commListRes.data.commissions.map((comm) => ({
      key: comm.date + commListRes.data.userId,
      date: comm.date,
      indirectRegistrationCount: comm.indirectRegistrationCount,
      directRegistrationCount: comm.directRegistrationCount,
      indirectRechargeAmount: comm.indirectRechargeAmount,
      directRechargeAmount: comm.directRechargeAmount,
    }));

    const now = new Date();
    if (ds.length == 0 || new Date(ds[0].date).getDate() < now.getDate()) {
      ds.unshift({
        key: commListRes.data.userId,
        date: now,
        indirectRegistrationCount: 0,
        directRegistrationCount: 0,
        indirectRechargeAmount: 0,
        directRechargeAmount: 0,
      });
    }
    setDataSource(ds);

    myfetch(`/v1/users/${userId}/commissions`).then((data) => {});
  }, [userId, commListRes.data]);

  if (userRes.error || commListRes.error) {
    return (userRes.error || commListRes.error).toString();
  }

  const columns = [
    {
      title: "日期",
      dataIndex: "date",
      key: "date",
      render: (value) => {
        return new Date(value).toLocaleDateString();
      },
    },
    {
      title: "注册量",
      dataIndex: "registration_count",
      key: "registration_count",
      render: (_, record) =>
        parseInt(record.indirectRegistrationCount) +
        parseInt(record.directRegistrationCount),
    },
    {
      title: "充值量",
      dataIndex: "recharge_amount",
      key: "recharge_amount",
      render: (_, record) =>
        `￥${
          (parseInt(record.indirectRechargeAmount) +
            parseInt(record.directRechargeAmount)) /
          100
        }`,
    },
    {
      title: "间接注册量",
      dataIndex: "indirectRegistrationCount",
      key: "indirect_registration_count",
      hidden: isDirectAgent,
    },
    {
      title: "直接注册量",
      dataIndex: "directRegistrationCount",
      key: "direct_registration_count",
      hidden: isDirectAgent,
    },
    {
      title: "间接充值量",
      dataIndex: "indirectRechargeAmount",
      key: "indirect_recharge_amount",
      render: (value) => parseInt(value) / 100,
      hidden: isDirectAgent,
    },
    {
      title: "直接充值量",
      dataIndex: "directRechargeAmount",
      key: "direct_recharge_amount",
      render: (value) => parseInt(value) / 100,
      hidden: isDirectAgent,
    },
  ];

  return (
    <Card title="推广明细">
      <Table
        columns={columns}
        dataSource={dataSource}
        loading={commListRes.isLoading || userRes.isLoading}
        size="small"
      />
    </Card>
  );
}
