"use client";
import { useRouter } from "next/navigation";
import React from "react";

export default function Main() {
  const router = useRouter();
  router.push("/dashboard");
  return;
}
