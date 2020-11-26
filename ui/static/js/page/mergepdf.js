FilePond.registerPlugin(
    FilePondPluginFileValidateType,
    FilePondPluginMediaPreview
);
const inputElement = document.querySelector('input[type="file"]');
const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/mergepdf',
    instantUpload: false,
    allowFileTypeValidation: true,
    acceptedFileTypes: ["application/pdf"],
    allowProcess: false,  // disable upload(up-arrow) icon on file
    allowReorder: true,
    labelIdle: '拖放文件或点击',
    labelFileTypeNotAllowed: '格式错误！',
    fileValidateTypeLabelExpectedTypes: '是PDF格式的文件吗？'
});

const mergeBtn = document.getElementById('upload');
var spinner = document.getElementById('spinner');

mergeBtn.addEventListener('click',function(){
    var files = pond.getFiles();
    if(files.length<=1){
        notify('请选择多个PDF文件进行合并！')
        return;
    }
    spinner.style.display = "block";
    this.hidden = true;

    myAjaxUpload(files,"merge");
});
