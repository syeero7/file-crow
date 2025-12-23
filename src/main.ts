import "./style.css";
import "./file_upload.ts";
import { connectWebsocket } from "./websocket";

connectWebsocket((s) => console.log(s));
