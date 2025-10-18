import {createApp} from "vue";
import App from "./App.vue";
import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import HomePage from "./pages/HomePage.vue";
import {Store} from "@tauri-apps/plugin-store";
import LoginPage from "./pages/LoginPage.vue";

export enum RouteNames {
    Home = "home",
    Login = "login",
}

const routes: RouteRecordRaw[] = [
    {path: "/home", component: HomePage, name: RouteNames.Home},
    {path: "/login", component: LoginPage, name: RouteNames.Login},
    {path: "/", redirect: "/home"},
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
});

let store: Store;

export async function refreshStore() {
    if (!store) {
        store = await Store.load("store.json");
    } else {
        await store.reload();
    }
}

router.beforeEach(async (to, _, next) => {
    const user = await store.get("user");

    if (user && to.name === RouteNames.Login) {
        return next({name: RouteNames.Home});
    }

    if (!user && to.name !== RouteNames.Login) {
        return next({name: RouteNames.Login});
    }

    next();
});

async function bootstrap() {
    await refreshStore();

    const app = createApp(App);

    app.provide("store", store);

    app.use(router)
    app.mount("#app");
}

void bootstrap();

