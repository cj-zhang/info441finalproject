import React, { Component } from "react";
import classnames from "classnames";

import "./style.css";

export default class Footer extends Component {
	render() {
		return(
			<footer className={classnames("Footer", this.props.className)}>
				<h1>Footer</h1>
				<p>TODO: REPLACE Footer <code>./components/Layout/Footer/index.js</code></p>
			</footer>
		);
	}
}
