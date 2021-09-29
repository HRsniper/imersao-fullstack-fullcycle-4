import { Container, makeStyles } from "@material-ui/core";
import { NextPage } from "next";
import { ReactNode } from "react";
import Navbar from "./Navbar";

const useStyles = makeStyles({
  container: {
    height: "calc(100% - 64px)"
  }
});

type PageProps = {
  children?: ReactNode | any;
};

export const Page: NextPage<PageProps> = (props) => {
  const classes = useStyles();

  return (
    <>
      <Navbar />
      <Container className={classes.container}>{props.children}</Container>
    </>
  );
};

export default Page;
