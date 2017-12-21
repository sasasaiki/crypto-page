$(function () {
	init()
})

function init() {
	let $msgBox = $("#chatbox textarea")
	let $messages = $("#messages")
	let $chatBox = $("#chatbox")

	$chatBox.submit(function (): boolean { onClickSubmit($msgBox, $messages); return false; })
}

function onClickSubmit($msgBox: JQuery, $messages: JQuery) {
	let msg = $msgBox.val().toString()
	alert(msg)
	$messages.append($("<li>").text(msg))
	$msgBox.val("")
}
