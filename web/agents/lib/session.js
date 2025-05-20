import useSWR from "swr";
import { NotLoggedInError } from "./error";

export function useUserId() {
  if (typeof window == "undefined") {
    return;
  }

  return localStorage.getItem("userId");
}

export function useToken() {
  if (typeof window == "undefined") {
    return;
  }

  return localStorage.getItem("token");
}

// 自动携带 token, 拼接完整路径
export async function myfetch(path, option) {
  const token = useToken();

  option = option || {};
  option.headers = {
    ...(option.headers || {}),
    Authorization: "Bearer " + token,
  };

  const res = await fetch(process.env.NEXT_PUBLIC_API_BASE + path, option);
  if (!res.ok) {
    throw new Error(await res.text());
  }
  return res.json();
}

export function useUser() {
  const userId = useUserId();

  const { data, isLoading, error } = useSWR(`/v1/users/${userId}`, myfetch);
  return { user: data, isLoading, error };
}

export function logOut() {
  if (typeof window == "undefined") {
    return;
  }

  localStorage.removeItem("token");
  localStorage.removeItem("userId");
}
