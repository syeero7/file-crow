import { getContext, setContext } from "svelte";
import { SvelteMap } from "svelte/reactivity";

import { toReadableSize, calculateProgress, type Transfer } from "@/lib/utils";
import { streamFile } from "@/lib/api";

const WS_KEY = "websocket";

class WS {
  #files = new SvelteMap<string, Transfer>();
  #ws: WebSocket | null = null;
  pendingFiles = new Map<string, File>();

  connect() {
    if (this.#ws != null) return;
    this.#ws = new WebSocket(`ws://${location.hostname}:8080/ws`);

    this.#ws.onopen = () => {
      console.info("websocket connected");
    };

    this.#ws.onclose = (e) => {
      console.error(`websocket disconnected code: ${e.code}`);
    };

    this.#ws.onmessage = (e) => {
      try {
        const msg = JSON.parse(e.data);
        if (!isWebSocketMsg(msg)) {
          console.error(`unexpected message type: ${typeof e.data}`);
          return;
        }

        switch (msg.type) {
          case "register": {
            this.#files.set(msg.id, {
              name: msg.name,
              status: "pending",
              size: toReadableSize(msg.size),
              progress: { percentage: 0 },
            });
            break;
          }

          case "progress": {
            const { current, total, id } = msg;
            const old = this.#files.get(id)!;
            const { speed, newProgress } = calculateProgress(old.progress, current, total);
            this.#files.set(id, {
              ...old,
              progress: newProgress,
              status: speed,
            });
            break;
          }

          case "ready": {
            const file = this.pendingFiles.get(msg.id);
            if (file) streamFile(msg.id, file).then(() => this.pendingFiles.delete(msg.id));
            break;
          }

          case "done":
          case "failed": {
            const old = this.#files.get(msg.id)!;
            this.#files.set(msg.id, { ...old, status: msg.type });
            setTimeout(() => {
              this.#files.delete(msg.id);
            }, 2000);
            break;
          }
        }
      } catch (err) {
        console.error(err);
      }
    };
  }

  get files(): MapIterator<[string, Transfer]> {
    return this.#files.entries();
  }
}

export function initializeWebsocket() {
  const ws = new WS();
  ws.connect();
  setContext(WS_KEY, ws);
  return ws;
}

export function useWebsocket() {
  return getContext<WS>(WS_KEY);
}

export type RegisterFileMsg = {
  type: "register";
  id: string;
  name: string;
  size: number;
};

type TransferProgressMsg = {
  type: "progress";
  id: string;
  current: number;
  total: number;
};

type FileTranferStateMsg = {
  type: "ready" | "done" | "failed";
  id: string;
};

type WebSocketMsg = RegisterFileMsg | TransferProgressMsg | FileTranferStateMsg;

function isWebSocketMsg(json: Record<string, unknown>): json is WebSocketMsg {
  return "id" in json && "type" in json;
}
