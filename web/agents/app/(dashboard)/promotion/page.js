import { Table } from "antd";

export default function Main() {
  const columns = [
    {
      title: "日期",
      dataIndex: "date",
      key: "date",
    },
    {
      title: "注册数量",
      dataIndex: "registrations",
      key: "registrations",
    },
    {
      title: "总佣金",
      dataIndex: "total_commission",
      key: "total_commission",
    },
  ];

  return <Table columns={columns} />;
}
