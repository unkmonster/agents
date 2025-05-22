"use client";
import { myfetch, useToken, useUserId } from "@/lib/session";
import { Button, Card, Modal, Space, Table } from "antd";
import { useState } from "react";
import { Input } from "antd";
import useSWR from "swr";

function DomainsTable({ domains, deleteDomain }) {
  const columns = [
    {
      title: "ID",
      dataIndex: "id",
      key: "id",
    },
    {
      title: "域名",
      dataIndex: "domain",
      key: "domain",
    },
    {
      title: "状态",
      dataIndex: "status",
      key: "status",
    },
    {
      title: "创建日期",
      dataIndex: "createdAt",
      key: "createdAt",
    },
    {
      title: "操作",
      dataIndex: "action",
      key: "action",
      render: (_, record) => {
        return (
          <Space size={"middle"}>
            <Button
              onClick={() => {
                deleteDomain(record.id);
              }}
              type="link"
            >
              删除
            </Button>
          </Space>
        );
      },
    },
  ];
  return (
    <Table
      columns={columns}
      dataSource={domains.map((domain) => ({
        key: domain.id,
        id: domain.id,
        domain: domain.domain,
      }))}
      size="middle"
    />
  );
}

export default function Main() {
  const userId = useUserId();
  const token = useToken();
  const [err, setErr] = useState();
  const [showModal, setShowModal] = useState(false);
  const [domain, setDomain] = useState("");
  const [confirmLoading, setConfirmLoading] = useState(false);

  const createDomain = async () => {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_BASE}/v1/users/${userId}/domains`,
      {
        method: "POST",
        headers: {
          "content-type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          domain,
        }),
      }
    );
    if (!res.ok) {
      setErr(await res.text());
    }
  };

  const deleteDomain = async (domain) => {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_BASE}/v1/domains/${domain}`,
      {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    if (!res.ok) {
      setErr(await res.text());
      return;
    }
    mutate();
  };

  const { data, isLoading, error, mutate } = useSWR(
    `/v1/users/${userId}/domains`,
    myfetch
  );

  return (
    <>
      <Modal
        title="添加域名"
        open={showModal}
        onCancel={() => setShowModal(false)}
        onOk={() => {
          setConfirmLoading(true);
          createDomain();
          setConfirmLoading(false);
          setShowModal(false);
          mutate();
        }}
        confirmLoading={confirmLoading}
      >
        <Input
          placeholder="输入域名"
          value={domain}
          onChange={(e) => {
            setDomain(e.target.value);
          }}
        />
      </Modal>

      <Card
        title="域名管理"
        extra={
          <Button type="primary" onClick={() => setShowModal(true)}>
            添加
          </Button>
        }
        loading={isLoading}
      >
        {error ? (
          JSON.stringify(error)
        ) : (
          <DomainsTable
            domains={data?.domains || []}
            deleteDomain={deleteDomain}
          />
        )}
      </Card>
    </>
  );
}
