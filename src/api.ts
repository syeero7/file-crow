import type { RegisterFileMsg } from "./websocket";

export function registerFile(body: RegisterFileMsg) {
  return fetcher("/register", "POST", body, ["json"]);
}

async function fetcher(
  path: string,
  method: "GET" | "POST",
  body?: Record<string, unknown>,
  headers?: ("json" | "multipart")[],
) {
  const base = import.meta.env.DEV ? "/api" : `/`;
  const url = `${base}${path}`;
  const options: RequestInit = { method };
  const tmp: Record<string, string> = {};

  if (body && headers) {
    headers.forEach((header) => {
      switch (header) {
        case "json": {
          tmp["Content-Type"] = "application/json";
          options.body = JSON.stringify(body);
          break;
        }

        case "multipart": {
          const formData = new FormData();
          for (const key in body) {
            const val = body[key];
            if (!(val instanceof Blob))
              throw new Error(`$${key} is not a Blob`);
            formData.append(key, val);
            options.body = formData;
          }
        }
      }
    });

    options.headers = tmp;
  }

  const res = await fetch(url, options);
  if (!res.ok) throw res;
}
