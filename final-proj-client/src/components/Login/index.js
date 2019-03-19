import React from "react";
import {Link} from "react-router-dom";


export default class SignInView extends React.Component {
    // constructor(props) {
    //     super(props);
    //     this.state = {
    //         errorMessage: undefined,
    //         currentUser: undefined,
    //         email: "",
    //         password: ""
    //     };
    // }
    // componentDidMount() {
    //     this.authUnsub = firebase.auth().onAuthStateChanged(user => {
    //       this.setState({currentUser: user});
    //     });
    // }
    // componentWillUnmount() {
    //     this.authUnsub();
    // }
    // handleSignIn(evt) {
    //     evt.preventDefault();
    //     firebase.auth().signInWithEmailAndPassword(this.state.email, this.state.password)
    //         .then(() => this.props.history.push("/channels/General"))
    //         .catch(err => alert(err.message));
    // }
       
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

                <form>
                    <div className="form-group">
                        <h4 style={labelStyle}>Email:</h4>
                        <input id="email" type="email" className="form-control" 
                        placeholder="enter your email address"/>
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Password:</h4>
                        <input id="password" type="password" className="form-control"
                        placeholder="enter your password" />
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