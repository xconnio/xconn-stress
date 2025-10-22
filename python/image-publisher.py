from time import sleep

import cv2
from xconn.client import connect_anonymous


def capture_image_to_bytes(camera_index: int = 0, image_format: str = ".jpg") -> bytes:
    cap = cv2.VideoCapture(camera_index)
    if not cap.isOpened():
        raise RuntimeError("Could not open camera.")

    ret, frame = cap.read()
    cap.release()

    if not ret:
        raise RuntimeError("Failed to capture image from camera.")

    success, buffer = cv2.imencode(image_format, frame)
    if not success:
        raise RuntimeError("Failed to encode image.")

    return buffer.tobytes()


session = connect_anonymous("ws://192.168.0.109:8080/ws", "realm1")
while True:
    img_bytes = capture_image_to_bytes()
    session.publish("io.xconn.image", [img_bytes])
    sleep(5)
