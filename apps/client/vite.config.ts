import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";

// @ts-expect-error process is a nodejs global
const host = process.env.TAURI_DEV_HOST;
// @ts-expect-error process is a nodejs global
const port = parseInt(process.env.TAURI_DEV_PORT || "1420");

// https://vite.dev/config/
export default defineConfig(async () => ({
    plugins: [vue()],

    // Vite options tailored for Tauri development and only applied in `tauri dev` or `tauri build`
    //
    // 1. prevent Vite from obscuring rust errors
    clearScreen: false,
    // 2. tauri expects a fixed port, fail if that port is not available
    server: {
        port,
        strictPort: false, // autorise un port alternatif si celui-ci est pris
        host: host || false,
        hmr: host
            ? {
                protocol: "ws",
                host,
                port: port + 1, // websocket HMR sur port différent
            }
            : undefined,
        watch: {
            ignored: ["**/src-tauri/**"],
        },
    },
}));
