"use client";
import { myfetch, useToken, useUserId } from "@/lib/session";
import {
  Alert,
  Button,
  Card,
  Modal,
  notification,
  Result,
  Space,
  Table,
} from "antd";
import { createContext, useState } from "react";
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
      size="small"
    />
  );
}

const Context = createContext({ name: "Default" });

export default function Main() {
  const [api, contextHolder] = notification.useNotification();

  const userId = useUserId();
  const [showModal, setShowModal] = useState(false);
  const [domain, setDomain] = useState("");
  const [confirmLoading, setConfirmLoading] = useState(false);

  const { data, isLoading, error, mutate } = useSWR(
    `/v1/users/${userId}/domains`,
    myfetch
  );

  const createDomain = async () => {
    try {
      await myfetch(`/v1/users/${userId}/domains`, {
        method: "POST",
        headers: {
          "content-type": "application/json",
        },
        body: JSON.stringify({
          domain,
        }),
      });
      mutate();
    } catch (err) {
      api.info({
        message: "创建失败",
        description: JSON.stringify(err),
        placement: "topRight",
      });
    }
  };

  const deleteDomain = async (domain) => {
    try {
      const res = await myfetch(`/v1/domains/${domain}`, {
        method: "DELETE",
      });
      if (!res.ok) {
        setErr(await res.text());
        return;
      }
      mutate();
    } catch (err) {
      api.info({
        message: "删除失败",
        description: JSON.stringify(err),
        placement: "topRight",
      });
    }
  };

  return (
    <Context.Provider value={null}>
      <>
        {contextHolder}
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
            <Result
              status={"error"}
              title={"加载失败"}
              subTitle={JSON.stringify(error)}
            />
          ) : (
            <DomainsTable
              domains={data?.domains || []}
              deleteDomain={deleteDomain}
            />
          )}
        </Card>
      </>
    </Context.Provider>
  );
}
