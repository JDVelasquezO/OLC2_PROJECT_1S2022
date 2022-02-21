myCodeMirror = CodeMirror.fromTextArea(document.getElementById("writer"), {
    lineNumbers: true,
    styleActivateLine: true,
    matchBrackets: true,
    theme: 'dracula',
    mode: 'javascript',
});

myCodeMirror2 = CodeMirror.fromTextArea(document.getElementById("reader"), {
    lineNumbers: true,
    theme: 'dracula',
    mode: 'powershell',
    readOnly: true,
});