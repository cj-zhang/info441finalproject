import React from "react";
import {Link} from "react-router-dom";


export default class SignInView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            errorMessage: undefined,
            currentUser: undefined,
            email: "",
            password: ""
        };
    }
    // componentDidMount() {
    //     this.authUnsub = firebase.auth().onAuthStateChanged(user => {
    //       this.setState({currentUser: user});
    //     });
    // }
    // componentWillUnmount() {
    //     this.authUnsub();
    // }
    // successfulAuth() {
    //     // return Promise.resolve(this.props.history.push("/"))
    //     console.log("successful signin");
    // }
    handleSignIn(evt) {
        evt.preventDefault();
        return fetch('http://smash.chenjosephzhang.me/v1/sessions', {
            method: "POST", 
            mode: "cors", // no-cors, cors, *same-origin
            cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
            credentials: "include", // include, *same-origin, omit
            headers: {
                "Content-Type": "application/json",
                // "Content-Type": "application/x-www-form-urlencoded",
            },
            redirect: "follow", // manual, *follow, error
            referrer: "no-referrer", // no-referrer, *client
            body: {
                "Email": this.state.email,
                "Password": this.state.password
            }, // body data type must match "Content-Type" header
        })
        .then(response => response.status())
        .then(function(response) {
            if(response === 201) {
                // successfulAuth();
                console.log("SUCCESSFUL");
            } else {
                alert("Unsuccessful log in attempt");
            }
        })
        .catch(function(error) {
            console.log('There has been a problem with your fetch operation: ', error.message);
        });
    }    

    render() {    
        let signinStyle = {
            width: "30%",
            marginTop: "100px",
            marginLeft: "auto",
            marginRight: "auto",
            backgroundColor: 'rgba(0,0,0,0.2)',
            borderRadius: '8px'
        }
        let titleStyle = {
            textAlign: "center",
            fontSize: 60,
            color: "white"
        }
        let labelStyle = {
            color: "white"
        }
        let textStyle = {
            marginLeft: "8px",
            paddingTop: "5px"
        }
        return (
            <section style={signinStyle}> 
                <div className="container">
                <h1 style={titleStyle}>Sign In</h1>

                <form onSubmit={evt => this.handleSignIn(evt)}>
                    <div className="form-group">
                        <h4 style={labelStyle}>Email:</h4>
                        <input id="email" type="email" className="form-control" 
                        placeholder="enter your email address"
                        onInput={evt => this.setState({email: evt.target.value})}/>
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Password:</h4>
                        <input id="password" type="password" className="form-control"
                        placeholder="enter your password"
                        onInput={evt => this.setState({password: evt.target.value})}/>
                    </div>
                    <div className="last-row d-flex">
                        <div className="form-group">
                            <button type="submit" className="btn btn-primary">
                                Sign In
                            </button>
                        </div>
                        <p  style={textStyle}>Don't yet have an account? <Link to={"/signup"}>Sign Up!</Link></p>
                    </div>
                </form>

                
            </div>
        </section>
        );
    }
}