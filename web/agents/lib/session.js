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
  const token = "Bearer " + useToken();

  option = option || {};
  option.headers = {
    ...(option.headers || {}),
    Authorization: token,
  };

  const res = await fetch(process.env.NEXT_PUBLIC_API_BASE + path, option);
  return res.json();
}
