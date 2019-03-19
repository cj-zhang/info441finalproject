import React, { Component } from "react";
import classnames from "classnames";

import Header from "./Header";
import Footer from "./Footer";

import "./style.css";

export default class Layout extends Component {
	render() {
		return(
			<div className="body">
				<Header/>
				<main className={classnames("Main", this.props.className)}>
					<div className="container">
						{ this.props.children }
					</div>
				</main>
				{/* <Footer/> */}
			</div>
		);
	}
}
