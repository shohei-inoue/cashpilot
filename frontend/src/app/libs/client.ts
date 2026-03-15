// 共通の認証クライアント関数 (server actions - cookies + JWT)

import { cookies } from "next/headers";
import axios, { AxiosInstance, AxiosError } from "axios";

/** Server Actions 用: コンテナ内では API_URL（backend サービス）を参照する */
const getBaseUrl = () =>
  process.env.API_URL ?? process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

/**
 * Server Actions / Server Components 向けの API クライアント
 * リクエストの Cookie（JWT）を自動で転送してバックエンドに認証付きでリクエストする
 */

export const createAuthClient = async (): Promise<AxiosInstance> => {
  const cookieStore = await cookies();
  const cookieHeader = cookieStore.toString();

  const authClient = axios.create({
    baseURL: getBaseUrl(),
    headers: {
      "Content-Type": "application/json",
      ...(cookieHeader && { Cookie: cookieHeader }),
    },
  });

  authClient.interceptors.response.use(
    async (response) => {
      // signup/login の Set-Cookie を Next.js の Cookie に転送
      const setCookie = response.headers["set-cookie"];
      if (setCookie && response.status >= 200 && response.status < 300) {
        const raw = Array.isArray(setCookie) ? setCookie[0] : setCookie;
        const match = typeof raw === "string" && raw.match(/token=([^;]+)/);
        if (match) {
          const store = await cookies();
          store.set("token", match[1], {
            httpOnly: true,
            secure: process.env.NODE_ENV === "production",
            sameSite: "lax",
            maxAge: 86400,
            path: "/",
          });
        }
      }
      return response;
    },
    (error: AxiosError) => {
      if (error.response?.status === 401) {
        console.error("Unauthorized:", error.response?.data);
        throw new Error("Unauthorized");
      }
      return Promise.reject(error);
    }
  );

  return authClient;
}
