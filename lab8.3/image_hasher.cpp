#include "image_hasher.h"
#include <openssl/sha.h>
#include <iomanip>
#include <sstream>
#include <fstream>

std::string ImageHasher::calculateSHA256(const cv::Mat& image) {
    // Конвертируем изображение в байтовый массив
    std::vector<uchar> buffer;
    cv::imencode(".png", image, buffer);
    
    // Вычисляем SHA-256
    unsigned char hash[SHA256_DIGEST_LENGTH];
    SHA256(buffer.data(), buffer.size(), hash);
    
    // Конвертируем в hex-строку
    std::stringstream ss;
    for(int i = 0; i < SHA256_DIGEST_LENGTH; i++) {
        ss << std::hex << std::setw(2) << std::setfill('0') << (int)hash[i];
    }
    
    return ss.str();
}

void ImageHasher::modifyPixel(cv::Mat& image) {
    // Изменяем синий канал первого пикселя
    if(image.cols > 0 && image.rows > 0) {
        cv::Vec3b& pixel = image.at<cv::Vec3b>(0, 0);
        pixel[0] = (pixel[0] + 1) % 256; // Изменяем синий канал
    }
}

void ImageHasher::saveResultsToFile(const std::string& originalHash, 
                                  const std::string& modifiedHash, 
                                  const std::string& filename) {
    std::ofstream outFile(filename);
    if(outFile) {
        outFile << "Original hash: " << originalHash << "\n";
        outFile << "Modified hash: " << modifiedHash << "\n";
        outFile << "Hashes match: " << (originalHash == modifiedHash ? "true" : "false") << "\n";
    }
}