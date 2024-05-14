// This function should be used to open a confirm dialog box.
// it expects to receive the element that was clicked.
function openConfirm(el) {
    let title = el.dataset.confirmTitle? el.dataset.confirmTitle : 'Are you sure?'
    let text = el.dataset.confirmBody? el.dataset.confirmBody : 'This action cannot be undone.'
    let cbutton = el.dataset.confirmButton? el.dataset.confirmButton : 'Delete'
    //Confirm title, 
    //Confirm body
    return Swal.fire({
        title: title, 
        text: text, 
        showCancelButton: true, 
        showCloseButton: true, 
        customClass:{
            container: 'top-0 left-0 right-0 bottom-0', 
            title: 'text-lg font-semibold leading-6 text-gray-900 mb-1 inline-block flex', 
            cancelButton: 'bg-white mt-3 rounded-md border border-gray-300 shadow-sm px-4 py-2 text-base font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 lg:text-base w-40 sm:mt-0 w-auto text-sm', 
            actions: 'container-actions justify-between w-full rounded mt-7', 
            confirmButton: `bg-red-600 rounded-md border border-transparent shadow-sm px-4 py-2 text-base font-medium text-white focus:outline-none lg:text-base w-40 sm:ml-3 w-auto text-sm`,
        }, 
        width: '390px',
        buttonsStyling: false,
        reverseButtons: true, 
        confirmButtonText: cbutton, 
        cancelButtonText: 'Cancel'
    })
}