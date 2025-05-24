"use client";
import { logOut, myfetch, useUserId } from "@/lib/session";
import { UserContext } from "@/lib/user_context";
import {
  Alert,
  Button,
  Card,
  Form,
  Input,
  message,
  Modal,
  Result,
  Space,
  Spin,
  Table,
} from "antd";
import useMessage from "antd/es/message/useMessage";
import { useRouter } from "next/navigation";
import { useContext, useEffect, useState } from "react";
import useSWR from "swr";

function CreateChildModal({ show, onClose, messageApi }) {
  const [loading, setLoading] = useState(false);
  const [err, setErr] = useState();
  const [form] = Form.useForm();
  const user = useContext(UserContext);
  const router = useRouter();

  function createChild(values) {
    return myfetch(`/v1/auth/register`, {
      method: "POST",
      body: JSON.stringify({
        username: values.username,
        password: values.password,
        nickname: values.nickname,
        parentId: user.id,
        level: user.level + 1,
        sharePercent: parseFloat(values.sharePercent),
      }),
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  const onFinish = async (values) => {
    setLoading(true);
    try {
      await createChild(values);
      onClose();
      messageApi.open({
        type: "success",
        content: "开户成功",
      });
    } catch (err) {
      if (err.code == 401) {
        logOut();
        router.push("/login");
      }
      setErr(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      title="添加下级代理"
      open={show}
      onCancel={onClose}
      onOk={() => {
        form.submit();
      }}
      confirmLoading={loading}
      afterClose={() => setErr(null)}
    >
      <Space>
        <Form size="small" onFinish={onFinish} form={form}>
          {err && (
            <Alert
              message={JSON.stringify(err)}
              type="error"
              style={{ margin: "15px 0" }}
            />
          )}
          <Form.Item
            label="用户名"
            name={"username"}
            rules={[{ required: true }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="密码"
            name={"password"}
            rules={[{ required: true }]}
          >
            <Input.Password />
          </Form.Item>
          <Form.Item
            label="分成比例"
            name={"sharePercent"}
            rules={[{ required: true }]}
          >
            <Input />
          </Form.Item>
        </Form>
      </Space>
    </Modal>
  );
}

export default function Main() {
  const columns = [
    {
      title: "ID",
      dataIndex: "id",
    },
    {
      title: "用户名",
      dataIndex: "username",
    },
    {
      title: "代理等级",
      dataIndex: "level",
    },
    {
      title: "分成比例",
      dataIndex: "sharePercent",
    },
    {
      title: "最后登录日期",
      dataIndex: "lastLoginAt",
    },
    {
      title: "注册日期",
      dataIndex: "created_at",
    },
  ];

  const user = useContext(UserContext);
  const childrenRes = useSWR(`/v1/users/${user?.id}/children`, myfetch);
  const router = useRouter();
  const [showModal, setShowModal] = useState(false);
  const [messageApi, messageHolder] = useMessage();
  useEffect(() => {
    if (childrenRes.error && childrenRes.error.isUnauthorized()) {
      logOut();
      router.push("/login");
    }
  }, [childrenRes.error]);

  return (
    <>
      {messageHolder}
      <CreateChildModal
        show={showModal}
        onClose={() => setShowModal(false)}
        messageApi={messageApi}
      />

      <Card
        title="下级管理"
        extra={
          <Button type="primary" onClick={() => setShowModal(true)}>
            开户
          </Button>
        }
      >
        {childrenRes.error ? (
          <Result title={JSON.stringify(childrenRes.error)} />
        ) : (
          <Table
            loading={childrenRes.isLoading}
            size="small"
            columns={columns}
            dataSource={
              childrenRes.data &&
              childrenRes.data.users.map((user) => ({
                key: user.id,
                id: user.id,
                username: user.username,
                created_at: user.createdAt,
                level: user.level,
                sharePercent: user.sharePercent,
                lastLoginAt: user.lastLoginAt,
              }))
            }
          />
        )}
      </Card>
    </>
  );
}
