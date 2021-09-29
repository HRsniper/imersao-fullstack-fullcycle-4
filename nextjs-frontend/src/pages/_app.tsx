import { CssBaseline, ThemeProvider } from "@material-ui/core";
import { AuthClientError, AuthClientEvent } from "@react-keycloak/core";
import { getKeycloakInstance, SSRCookies, SSRKeycloakProvider } from "@react-keycloak/ssr";
import type { AppContext, AppProps } from "next/app";
import Head from "next/head";
import { useEffect } from "react";
import { TenantProvider } from "../contexts/TenantProvider";
import "../styles/globals.css";
import { KEYCLOAK_PUBLIC_CONFIG } from "../utils/auth";
import { parseCookies } from "../utils/cookies";
import { keycloakEvents$ } from "../utils/http";
import { theme } from "../utils/theme";

function MyApp({ Component, pageProps, cookies }: AppProps & { cookies: any }) {
  useEffect(() => {
    const jssStyles = document.querySelector("#jss-server-side");
    jssStyles?.parentElement?.removeChild(jssStyles);
  }, []);

  const onEvent = async (event: AuthClientEvent, error?: AuthClientError | undefined) => {
    if (event === "onAuthSuccess") {
      keycloakEvents$.next({
        type: "success"
      });
    }

    if (event === "onAuthError") {
      keycloakEvents$.next({
        type: "error"
      });
    }

    if (event === "onTokenExpired") {
      console.log("onTokenExpired");
      await getKeycloakInstance(null as any).updateToken(30);
    }
  };

  return (
    <SSRKeycloakProvider
      keycloakConfig={KEYCLOAK_PUBLIC_CONFIG}
      persistor={SSRCookies(cookies)}
      initOptions={{
        onLoad: "check-sso",
        silentCheckSsoRedirectUri:
          typeof window !== "undefined" ? `${window.location.origin}/silent-check-sso.html` : null
      }}
      onEvent={onEvent}
    >
      <TenantProvider>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Head>
            <meta name="viewport" content="width=device-width, initial-scale=1.0" />
          </Head>
          <Component {...pageProps} />
        </ThemeProvider>
      </TenantProvider>
    </SSRKeycloakProvider>
  );
}

MyApp.getInitialProps = async (appContext: AppContext) => {
  return {
    cookies: parseCookies(appContext.ctx.req)
  };
};

export default MyApp;
