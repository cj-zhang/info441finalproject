import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Routes from "./routes";

import "./index.css";

const outlet = document.getElementById("root");
ReactDOM.render(<Router history={BrowserRouter} routes={Routes} />, outlet);
