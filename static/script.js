const fileInput = document.getElementById('fileInput');
const uploadButton = document.getElementById('uploadButton');
const downloadButton = document.getElementById('downloadButton');
let downloadUrl = '';

uploadButton.addEventListener('click', async () => {
    if (!fileInput.files.length) {
        alert('Please select an Excel file to upload');
        return;
    }

    const formData = new FormData();
    formData.append('file', fileInput.files[0]);

    try {
        const response = await fetch('/api/translate', {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            throw new Error(`Error ${response.status}: ${await response.text()}`);
        }

        downloadUrl = URL.createObjectURL(await response.blob());
        downloadButton.disabled = false;
    } catch (error) {
        alert(`Failed to upload and translate file: ${error.message}`);
    }
});

downloadButton.addEventListener('click', () => {
    if (!downloadUrl) {
        alert('No file available for download');
        return;
    }

    const link = document.createElement('a');
    link.href = downloadUrl;
    link.download = 'translated.xlsx';
    link.click();
});

