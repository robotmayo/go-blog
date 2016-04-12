"use strict";
import React from 'react';

export default class NewPost extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      postTitle : '',
      postBody : ''
    };
  }
  
  updateInput(key, e){
    console.log(e.target.value);
    const newState = Object.assign({}, this.state);
    newState[key] = e.target.value;
    this.setState(newState);
  }
  
  createPost(){
    return fetch('/api/new/post', {
            method: 'POST',
            headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              title : this.state.postTitle,
              body : this.state.postBody
            })
          })
          .then(response => response.json())
          .then(response => console.log(response))
          .catch(err => console.log(err));
  }
  
  render(){
    return (
      <div>
        <h3>Title</h3>
        <input type="text" value={this.state.postTitle} onChange={this.updateInput.bind(this, "postTitle")}/>
        <h3>Body</h3>
        <input type="text" value={this.state.postBody} onChange={this.updateInput.bind(this, "postBody")}/>
        <br />
        <button onClick={this.createPost.bind(this)}>Submit</button>
      </div>
    );
  }
  
}

