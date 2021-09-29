import { isEqual } from "lodash";
import React, { createContext, useEffect, useState } from "react";
import { useAuthSwr } from "../hooks/useAuthSwr";

interface Tenant {
  id: string;
  name: string;
  subdomain: string;
  balance: number;
}

const TenantContext = createContext<Tenant>(null as any);

const TenantProvider: React.FunctionComponent = (props) => {
  const [tenant, setTenant] = useState<Tenant>();
  const { data, error } = useAuthSwr("my-account", {
    refreshInterval: 10000
  });

  useEffect(() => {
    if (!isEqual(data, tenant)) {
      setTenant(data);
    }
  }, [data, tenant]);

  return <TenantContext.Provider value={tenant as any}>{props.children}</TenantContext.Provider>;
};

export default TenantContext;
export { TenantProvider };
export type { Tenant };
