from flask import Flask, request, make_response, jsonify, send_file #, render_template
import logging, os
import pandas as pd
import uuid


import onnxruntime
import numpy as np
import os

from skimage.transform import resize
import skimage.io

app = Flask(__name__, static_url_path='', static_folder = "uploads")
# file_handler = logging.FileHandler('server.log')
# app.logger.addHandler(file_handler)
app.logger.setLevel(logging.INFO)
PROJECT_HOME = os.path.dirname(os.path.realpath(__file__))
UPLOAD_FOLDER = '{}/uploads/'.format(PROJECT_HOME)
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

def create_new_folder(local_dir):
    newpath = local_dir
    if not os.path.exists(newpath):
        os.makedirs(newpath)
    return newpath

@app.route('/', methods = ['POST'])
def api_root():
    app.logger.info(PROJECT_HOME)
    app.logger.info('Get a request!')
    if request.method == 'POST' and request.get_json(): # request.files['image_file']:# and request.files['private_key_file']:
        app.logger.info(app.config['UPLOAD_FOLDER'])
        data = request.get_json()
        img_url = data['img_url']
        # image_file = request.files['image_file']
        image_file = os.path.join(app.config['UPLOAD_FOLDER'], data['img_url'])
        ####################################################################
        ## ONNX file decryption

        # create_new_folder(app.config['UPLOAD_FOLDER'])
        # saved_path =
        ####################################################################
        # df = pd.read_csv(image_file)
        # df['cap'] = 8.5


        ####################################################################
        # m = Prophet(growth='logistic')
        # m.fit(df)
        # future = m.make_future_dataframe(periods=1826)
        # future['cap'] = 8.5
        # fcst = m.predict(future)
        # fig = m.plot(fcst)

        # os.system("wget https://drive.google.com/file/d/1e5rCysXbl8AS4ALONTZa579FPWp29eZT/view?usp=sharing")

        img_name = str(uuid.uuid1()) + '.jpg'
        # ori_saved_path = os.path.join(app.config['UPLOAD_FOLDER'], img_name + '.png')
        res_saved_path = os.path.join(app.config['UPLOAD_FOLDER'], img_name)
        # image_file.save(ori_saved_path)
        session = onnxruntime.InferenceSession("best_enas.onnx")

        # input_ = np.random.rand(1,3,256,256).astype('float32')
        input_ = skimage.io.imread(image_file)
        if len(input_.shape) == 3:
            img_resized = resize(input_, (256, 256, 3), anti_aliasing=True)
            input_ = np.expand_dims(img_resized.transpose(2,0,1), axis=0).astype(np.float32)
        elif len(input_.shape) == 4:
            pass
        print(input_.shape)
        results = session.run([], {"input1": input_})[0] # 1*3*256*256
        results = results[0,3:6,:,:].transpose(1,2,0)
        print(results.shape)
        
        ####################################################################



        skimage.io.imsave(res_saved_path, results)
        result = {'url': img_name, 'status': 20000}
        response = make_response(jsonify(result))
        response.headers['Content-Type'] = 'application/json'
        return response
    else:
        return "Where is the csv?"
#
#
# @app.route('/show', methods = ['GET'])
# def show():
#     images = os.listdir('/home/yongli/Documents/flask-api-upload-image-master/uploads')
#     return render_template('images.html', images=images)
#
@app.route('/<filename>', methods = ['GET'])
def get_image(filename):
    img_name = filename+'.png'
    dir_list = os.listdir(app.config['UPLOAD_FOLDER'])
    print(dir_list)
    if img_name in dir_list:
        return send_file(os.path.join(app.config['UPLOAD_FOLDER'], img_name), mimetype='image/jpeg')
    else:
        return 'No such a image named \"'+filename+'\"'










# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    create_new_folder(app.config['UPLOAD_FOLDER'])
    app.run(debug=False, host='0.0.0.0', port=5000)

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
