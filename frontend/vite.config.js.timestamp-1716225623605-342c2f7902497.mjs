// vite.config.js
import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "file:///F:/GolandProjects/WorkSpace/src/goProjects/yyblog/frontend/node_modules/vite/dist/node/index.js";
import vue from "file:///F:/GolandProjects/WorkSpace/src/goProjects/yyblog/frontend/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import vueJsx from "file:///F:/GolandProjects/WorkSpace/src/goProjects/yyblog/frontend/node_modules/@vitejs/plugin-vue-jsx/dist/index.mjs";
var __vite_injected_original_import_meta_url = "file:///F:/GolandProjects/WorkSpace/src/goProjects/yyblog/frontend/vite.config.js";
var vite_config_default = defineConfig({
  plugins: [
    vue(),
    vueJsx()
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url))
    }
  },
  server: {
    // host: '192.168.0.3',
    proxy: {
      "/uploads": {
        target: "http://154.26.129.43:9081",
        changeOrigin: true,
        pathRewrite: { "^/uploads": "/uploads" }
      },
      "/yyblog/api/v1": {
        target: "http://154.26.129.43:9081",
        changeOrigin: true,
        pathRewrite: { "^/yyblog/api/v1": "/yyblog/api/v1" }
      }
    },
    port: 80
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJGOlxcXFxHb2xhbmRQcm9qZWN0c1xcXFxXb3JrU3BhY2VcXFxcc3JjXFxcXGdvUHJvamVjdHNcXFxceXlibG9nXFxcXGZyb250ZW5kXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJGOlxcXFxHb2xhbmRQcm9qZWN0c1xcXFxXb3JrU3BhY2VcXFxcc3JjXFxcXGdvUHJvamVjdHNcXFxceXlibG9nXFxcXGZyb250ZW5kXFxcXHZpdGUuY29uZmlnLmpzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9GOi9Hb2xhbmRQcm9qZWN0cy9Xb3JrU3BhY2Uvc3JjL2dvUHJvamVjdHMveXlibG9nL2Zyb250ZW5kL3ZpdGUuY29uZmlnLmpzXCI7aW1wb3J0IHsgZmlsZVVSTFRvUGF0aCwgVVJMIH0gZnJvbSAnbm9kZTp1cmwnXHJcblxyXG5pbXBvcnQgeyBkZWZpbmVDb25maWcgfSBmcm9tICd2aXRlJ1xyXG5pbXBvcnQgdnVlIGZyb20gJ0B2aXRlanMvcGx1Z2luLXZ1ZSdcclxuaW1wb3J0IHZ1ZUpzeCBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUtanN4J1xyXG5cclxuLy8gaHR0cHM6Ly92aXRlanMuZGV2L2NvbmZpZy9cclxuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKHtcclxuICBwbHVnaW5zOiBbXHJcbiAgICB2dWUoKSxcclxuICAgIHZ1ZUpzeCgpLFxyXG4gIF0sXHJcbiAgcmVzb2x2ZToge1xyXG4gICAgYWxpYXM6IHtcclxuICAgICAgJ0AnOiBmaWxlVVJMVG9QYXRoKG5ldyBVUkwoJy4vc3JjJywgaW1wb3J0Lm1ldGEudXJsKSlcclxuICAgIH1cclxuICB9LFxyXG4gIHNlcnZlcjoge1xyXG4gICAgLy8gaG9zdDogJzE5Mi4xNjguMC4zJyxcclxuICAgIHByb3h5OiB7XHJcbiAgICAgICcvdXBsb2Fkcyc6IHtcclxuICAgICAgICB0YXJnZXQ6ICdodHRwOi8vMTU0LjI2LjEyOS40Mzo5MDgxJyxcclxuICAgICAgICBjaGFuZ2VPcmlnaW46IHRydWUsXHJcbiAgICAgICAgcGF0aFJld3JpdGU6IHsgJ14vdXBsb2Fkcyc6ICcvdXBsb2FkcycgfVxyXG4gICAgICB9LFxyXG4gICAgICAnL3l5YmxvZy9hcGkvdjEnOiB7XHJcbiAgICAgICAgdGFyZ2V0OiAnaHR0cDovLzE1NC4yNi4xMjkuNDM6OTA4MScsXHJcbiAgICAgICAgY2hhbmdlT3JpZ2luOiB0cnVlLFxyXG4gICAgICAgIHBhdGhSZXdyaXRlOiB7ICdeL3l5YmxvZy9hcGkvdjEnOiAnL3l5YmxvZy9hcGkvdjEnIH1cclxuICAgICAgfVxyXG4gICAgfSxcclxuICAgIHBvcnQ6IDgwLFxyXG4gIH1cclxufSlcclxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUE4VyxTQUFTLGVBQWUsV0FBVztBQUVqWixTQUFTLG9CQUFvQjtBQUM3QixPQUFPLFNBQVM7QUFDaEIsT0FBTyxZQUFZO0FBSnNOLElBQU0sMkNBQTJDO0FBTzFSLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQzFCLFNBQVM7QUFBQSxJQUNQLElBQUk7QUFBQSxJQUNKLE9BQU87QUFBQSxFQUNUO0FBQUEsRUFDQSxTQUFTO0FBQUEsSUFDUCxPQUFPO0FBQUEsTUFDTCxLQUFLLGNBQWMsSUFBSSxJQUFJLFNBQVMsd0NBQWUsQ0FBQztBQUFBLElBQ3REO0FBQUEsRUFDRjtBQUFBLEVBQ0EsUUFBUTtBQUFBO0FBQUEsSUFFTixPQUFPO0FBQUEsTUFDTCxZQUFZO0FBQUEsUUFDVixRQUFRO0FBQUEsUUFDUixjQUFjO0FBQUEsUUFDZCxhQUFhLEVBQUUsYUFBYSxXQUFXO0FBQUEsTUFDekM7QUFBQSxNQUNBLGtCQUFrQjtBQUFBLFFBQ2hCLFFBQVE7QUFBQSxRQUNSLGNBQWM7QUFBQSxRQUNkLGFBQWEsRUFBRSxtQkFBbUIsaUJBQWlCO0FBQUEsTUFDckQ7QUFBQSxJQUNGO0FBQUEsSUFDQSxNQUFNO0FBQUEsRUFDUjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
