import React from "react";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";

import ReactDOM from "react-dom/client";

import App from "./App";
import "./index.css";
import { QueryClient, QueryClientProvider } from "react-query";

const colors = {
  brand: {
    900: "#1a365d",
    800: "#153e75",
    700: "#2a69ac",
  },
};
const theme = extendTheme({ colors });

const queryClient = new QueryClient();
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ChakraProvider theme={theme}>
      <QueryClientProvider client={queryClient}>
        <App />
      </QueryClientProvider>
    </ChakraProvider>
  </React.StrictMode>
);
