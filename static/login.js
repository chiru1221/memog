'use strict';

const e = React.createElement;

class LikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { liked: false };
  }

  render() {
    return (
        <Button variant="contained" color="primary">
          Hello World
        </Button>
      );
  }
}

const domContainer = document.querySelector('#app');
ReactDOM.render(e(LikeButton), domContainer);
