import React from 'react';
import {ListItem} from 'material-ui/List';
import IconButton from 'material-ui/IconButton';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import BoardActions from './actions/BoardActions';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import LabelIcon from 'material-ui/svg-icons/action/label';
import LabelOutlineIcon from 'material-ui/svg-icons/action/label-outline';
import {grey400} from 'material-ui/styles/colors';

export default class StoryCard extends React.Component {
  constructor(props) {
    super(props);
    this.state={};
  }

  onClickHandler = () => {
    BoardActions.toggleTasksForStoryId(this.props.story_id);
  }

  onClickIconHandler = (event) => {
    event.stopPropagation();
    BoardActions.openStoryShowDialog(this.props.story_id);
  }


  // This is needed for testing the a component without the whole app context
  static childContextTypes = {
    muiTheme: React.PropTypes.object
  }

  getChildContext() {
    return {
      muiTheme: getMuiTheme()
    };
  }

  render() {
    var leftIcon = <LabelOutlineIcon color={this.props.color} />;
    if (this.props.active) {
      leftIcon = <LabelIcon color={this.props.color}/>;
    }
    var rightIcon = (<IconButton tooltip="Show story" onClick={this.onClickIconHandler}>
        <MoreVertIcon color={grey400} />
      </IconButton>);

    return <ListItem
      primaryText={this.props.name}
      onClick={this.onClickHandler}
      leftIcon={leftIcon}
      rightIconButton={rightIcon}
    />;
  }

}
