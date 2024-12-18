import { socket } from './socket.js';
import { setChannel, getChannel } from './localstore.js'


// channel functionality handler...
let channelForm = document.getElementById('channelForm')

channelForm.onsubmit = function (event) {
    event.preventDefault()

    const formData = new FormData(event.target)
    const channelName = formData.get("channel")
    const username = formData.get("username")
    console.log("channel name and username:", channelName, username)
    setChannel(channelName, username)
    var channelDetails = {
        "type":"channel",
        "channelName": channelName,
        "username": username

    }
    var strMessage = JSON.stringify(channelDetails)
    socket.send(strMessage)
}

const channel = document.getElementById("channel")
const username = document.getElementById("username")
channel.value = localStorage.getItem("channel")
username.value = localStorage.getItem("userName")

// Text message handler...

let messageForm = document.getElementById("messageForm")

messageForm.onsubmit = function (event) {
    const formData = new FormData(event.target)
    const message = formData.get("message")
    var messageDetails ={
        "type":"message",
        "channelName":localStorage.getItem('channel'),
        "username":localStorage.getItem("username"),
        "message":message
    }
    var strMessage = JSON.stringify(messageDetails)
    socket.send(strMessage)
}