FilePond.registerPlugin(
    FilePondPluginFileValidateType,
    FilePondPluginMediaPreview
);
const inputElement = document.querySelector('input[type="file"]');
const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/rotatepdf',
    instantUpload: false,
    allowFileTypeValidation: true,
    acceptedFileTypes: ["application/pdf"],
    allowProcess: false,  // disable upload(up-arrow) icon on file
    labelIdle: '拖放文件或点击',
    labelFileTypeNotAllowed: '格式错误！',
    fileValidateTypeLabelExpectedTypes: '是PDF格式的文件吗？'
});

const rotateBtn = document.getElementById('upload');
var spinner = document.getElementById('spinner');

rotateBtn.addEventListener('click',function(){
    var files = pond.getFiles();
    if(files.length==0){
        notify('请选择PDF格式的文件！');
        return;
    }

    var clockwise = document.getElementById('cw').checked;
    var counterclockwise = document.getElementById('ccw').checked;
    var degrees = document.getElementById('degrees').value;
    
    if (!clockwise&&!counterclockwise){
        notify("顺时针还是逆时针?");
        return;
    }

    var obj = {
        action: 'rotate',
        cw: clockwise,
        ccw: counterclockwise,
        degree: degrees
    }
    
    spinner.style.display = "block";
    this.hidden = true;

    myAjaxUpload(files,obj);
});
