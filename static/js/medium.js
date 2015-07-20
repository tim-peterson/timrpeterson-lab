/*var ContentEditable = require("react-contenteditable");
  var MyComponent = React.createClass({
    getInitialState: function(){
      return {html: "<b>Hello <i>World</i></b>"};
    },

    handleChange: function(evt){
      this.setState({html: evt.target.value});
    },

    render: function(){
      return <ContentEditable html={this.state.html} onChange={this.handleChange}/>
    }
  });*/

  //console.log("medium.js");




  var Input = React.createClass({
    getInitialState: function() {
      return {idea: ''};
    },
    handleSubmit: function(e) {

      e.preventDefault();

      var idea = this.state.idea.trim();


      if(!idea) return;

      this.setState({
        idea: '',
        message: 'Please wait...'
      });

      // I'm adding a callback to the form submit handler, so you can
      // keep all the state changes in the component.
      this.props.onFormSubmit({
        idea: idea,
      //  userEmail: userEmail,
        url: "/submittedIdea"
      }, function(data) {
        this.setState({ message: data });
      });
    },
    changeValue: function(e) {
      this.setState({
        name: e.target.value
      });
    },
    handleChange: function(event) {
      this.setState({idea: event.target.value});
    },
    render: function () {
      var idea = this.state.idea;
      return (
        <div className="row">
            <form className="col s6 offset-s3 formElem" onSubmit={ this.handleSubmit }>
              <div className="row">
                <div className="input-field col s12">
                <label for="textarea1">tell me something new...</label>
                  <textarea id="textarea1"  onChange={this.handleChange} name="idea" value={ this.state.idea } className="materialize-textarea">{idea}</textarea>
                  <button className="btn waves-effect waves-light z-depth-0 blue-grey darken-1" type="submit" value="submit" name="action">
                      <i className="material-icons">done</i>
                    </button>
                  <div className="result">{ this.state.message }</div>
                </div>
              </div>
            </form>
        </div>
      );
    }
  });

  //React.render(<Input/>, document.getElementById("main"));



  var FormComp = React.createClass({

    // To get rid of those input refs I'm moving those values
    // and the form message into the state
    getInitialState: function() {
      return {
        name: '',
        email: '',
        message: ''
      };
    },

    handleSubmit: function(e) {

      e.preventDefault();

      var userName = this.state.name.trim();
      var userEmail = this.state.email.trim();

      if(!userName || !userEmail) return;

      this.setState({
        name: '',
        email: '',
        message: 'Please wait...'
      });

      // I'm adding a callback to the form submit handler, so you can
      // keep all the state changes in the component.
      this.props.onFormSubmit({
        userName: userName,
        userEmail: userEmail,
        url: "/submittedIdea"
      }, function(data) {
        this.setState({ message: data });
      });
    },

    changeName: function(e) {
      this.setState({
        name: e.target.value
      });
    },

    changeEmail: function(e) {
      this.setState({
        email: e.target.value
      });
    },

    render: function() {
      // the message and the input values are all component state now
      return (
        <div>
          <div className="result">{ this.state.message }</div>
          <form className="formElem" onSubmit={ this.handleSubmit } method="post" action="/submittedIdea">
            Name: <input type="text" className="userName" name="userName" value={ this.state.name } onChange={ this.changeName } /><br/>
            Email: <input type="text" className="userEmail" name="userEmail" value={ this.state.email } onChange={ this.changeEmail } /><br/>
            <input type="submit" value="Submit" />
          </form >
        </div>
      );
    }
  });


  var RC= React.createClass({

    onFormSubmit: function(data, callback){

       $.ajax({
          url: this.props.url,
          dataType: 'json',
          type: 'POST',
          data: data,
          success: callback,
          error: function(xhr, status, err) {
            console.log(xhr);
          //  console.log(this.props.url, status, err.toString());
            console.error(this.props.url, status, err.toString());

          }.bind(this)
        });
    },

    render: function() {
      return <Input onFormSubmit={this.onFormSubmit} />
      //return <FormComp onFormSubmit={this.onFormSubmit} />
    }
  });

  //
  React.render(
    <RC/>,
    document.getElementById('idea')
  );


  /*React.render(
    <h1>Hello, world!</h1>,
    document.getElementById('example')
  );*/


/*"use strict";

  class Hello extends React.Component {
  render() {
    return <div>Hello {this.props.name}</div>;
  }
}

React.render(<Hello name="World" />, document.getElementById('container'));*/
