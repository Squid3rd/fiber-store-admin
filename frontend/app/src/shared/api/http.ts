import { env } from "../../app/config/env";

export type ApiError = {
  status: number;
  message: string;
  details: unknown;
};

async function readJsonSafely(res: Response) {
  const text = await res.text();

  if (!text) {
    return undefined;
  }
  try {
    return JSON.parse(text);
  } catch (error) {
    return error;
  }
}

export async function http<T>(path: string, init?: RequestInit): Promise<T> {
  const url = `${env.apiBaseUrl}${path}`;

  const res = await fetch(url, {
    ...init,
    headers: {
      ...(init?.body ? { "Content-Type": "application/json" } : {}),
      ...(init?.headers ?? {}),
    },
  });

  const body = await readJsonSafely(res);

  if (!res.ok) {
    const message =
      typeof body === "object" && body && "error" in body
        ? String(body.error)
        : `Request Failed: ${res.statusText}`;

    throw { status: res.status, message, details: body } satisfies ApiError;
  }

  return body as T;
}
