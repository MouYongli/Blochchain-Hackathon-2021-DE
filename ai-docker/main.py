from flask import Flask, request, make_response, jsonify, send_file #, render_template
import logging, os
import pandas as pd
from fbprophet import Prophet
import uuid

app = Flask(__name__, static_url_path='', static_folder = "uploads")
file_handler = logging.FileHandler('server.log')
app.logger.addHandler(file_handler)
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
    if request.method == 'POST' and request.files['csv_file'] and request.files['private_key_file']:
        app.logger.info(app.config['UPLOAD_FOLDER'])
        csv_file = request.files['csv_file']
        private_key_file = request.files['private_key_file']
        ####################################################################
        ## ONNX file decryption
        ####################################################################
        df = pd.read_csv(csv_file)
        df['cap'] = 8.5
        ####################################################################
        m = Prophet(growth='logistic')
        m.fit(df)
        future = m.make_future_dataframe(periods=1826)
        future['cap'] = 8.5
        fcst = m.predict(future)
        fig = m.plot(fcst)
        ####################################################################
        create_new_folder(app.config['UPLOAD_FOLDER'])
        img_name = str(uuid.uuid1())
        saved_path = os.path.join(app.config['UPLOAD_FOLDER'], img_name+'.png')
        fig.savefig(saved_path)
        result = {'url': img_name, 'status': 200}
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
    app.run(debug=False, host='0.0.0.0', port=5000)

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
