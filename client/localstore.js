export function setChannel(channel, userName){
    localStorage.setItem("channel", channel)
    localStorage.setItem("userName", userName)
}

export function getChannel(){
    console.log("ll", localStorage.getItem("channel"))
    return [localStorage.getItem("channel"), localStorage.getItem("userName")]
}


