import cookie from "cookie";
import Cookies from "js-cookie";

function parseCookies(req?: any) {
  if (!req || !req.headers) {
    return {};
  }

  return cookie.parse(req.headers.cookie || "");
}

function destroyCookie(key: string) {
  Cookies.remove(key);
}

export { parseCookies, destroyCookie };
