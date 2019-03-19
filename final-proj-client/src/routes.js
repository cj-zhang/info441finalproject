import React from "react";
import { Route } from 'react-router';
import Layout from "./components/Layout";
import Home from "./components/Home";
import Profile from "./components/Profile";
import Players from "./components/Players";

const Routes = (
  <Route path="/" component={Layout}>
    <Route component={Home} />
    <Route path="/profile" component={Profile} />
    <Route path="/players" component={Players} />
  </Route>
);

export default Routes;
