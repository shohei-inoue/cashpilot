import axios, { AxiosResponse } from "axios";

/**
 * APIレスポンスを処理し、ステータスコードに応じて例外をスローする
 * 成功時は response.data を返す
 */
export async function handleResponse<T>(
  promise: Promise<AxiosResponse<T>>
): Promise<T> {
  try {
    const response = await promise;
    return response.data;
  } catch (error: unknown) {
    if (axios.isAxiosError(error)) {
      const status = error.response?.status;
      const body = error.response?.data as { error?: { message?: string } } | undefined;
      const message = body?.error?.message;

      console.error("API request error:", error.message);
      console.error("Response status:", status);
      console.error("Error details:", message || body || error.toString());

      throw new Error(
        typeof message === "string" ? message : "APIリクエストエラーが発生しました"
      );
    } else if (error instanceof Error) {
      console.error("Unexpected error:", error.message);
      throw error;
    } else {
      console.error("Unknown error:", error);
      throw new Error("Unknown error");
    }
  }
}
