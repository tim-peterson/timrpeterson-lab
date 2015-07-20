
var Navbar = React.createClass({
  render: function() {
    return (
      <nav className="" role="navigation">
       <div className="nav-wrapper container"><a id="logo-container" href="#" className="brand-logo">Logo</a>
         <ul className="right hide-on-med-and-down">
           <li><a href="/about">About</a></li>
         </ul>

         <ul id="nav-mobile" className="side-nav">
           <li><a href="/team">team</a></li>
         </ul>
         <a href="#" data-activates="nav-mobile" className="button-collapse"><i className="material-icons">Sign in</i></a>
       </div>
     </nav>
   );
 }
});

React.render(<Navbar />, document.getElementById("Navbar"));


var Footer = React.createClass({
  render: function(){
    return (
      <footer className="page-footer">
        <div className="container">
          <div className="row">
            <div className="col l6 s12">
              <p className="grey-text">Being active is a great way to stimulate your creativity. <a href="http://news.stanford.edu/news/2014/april/walking-vs-sitting-042414.html" target="_blank">Research shows</a> that in particular, walking, helps create new ideas. We created this site to help people share their ideas while walking. </p>
            </div>
            <div className="col l3 s12">
              <h5 className="grey-text"></h5>
              <ul>
                <li><a className="grey-text" href="https://facebook.com/ideaswhilewalking">Facebook</a></li>
                <li><a className="grey-text" href="https://twitter.com/ideaswhilewalking">Twitter</a></li>

              </ul>
            </div>
          </div>
        </div>
        <div className="footer-copyright">
          <div className="container">
          Made by <a className="orange-text text-lighten-3" href="http://timrpeterson.com">timrpeterson</a>
          </div>
        </div>
      </footer>
    )
  }
});

React.render(<Footer />, document.getElementById("footer"));

console.log("testing");

$('.button-collapse').sideNav({
     menuWidth: 300, // Default is 240
     edge: 'right', // Choose the horizontal origin
     closeOnClick: true // Closes side-nav on <a> clicks, useful for Angular/Meteor
   }
 );
