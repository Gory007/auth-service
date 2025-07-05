#include "image_hasher.h"
#include <opencv2/opencv.hpp>
#include <iostream>

int main() {
    // Загрузка изображения
    std::string imagePath;
    std::cout << "Enter image path: ";
    std::cin >> imagePath;
    
    cv::Mat image = cv::imread(imagePath, cv::IMREAD_COLOR);
    if(image.empty()) {
        std::cerr << "Could not open or find the image!" << std::endl;
        return -1;
    }
    
    // Вычисляем оригинальный хеш
    std::string originalHash = ImageHasher::calculateSHA256(image);
    std::cout << "Original hash: " << originalHash << std::endl;
    
    // Модифицируем пиксель
    ImageHasher::modifyPixel(image);
    
    // Вычисляем хеш после модификации
    std::string modifiedHash = ImageHasher::calculateSHA256(image);
    std::cout << "Modified hash: " << modifiedHash << std::endl;
    
    // Сравниваем хеши
    bool hashesMatch = (originalHash == modifiedHash);
    std::cout << "Hashes match: " << (hashesMatch ? "true" : "false") << std::endl;
    
    // Сохраняем результаты в файл
    ImageHasher::saveResultsToFile(originalHash, modifiedHash, "output/hashes.txt");
    
    return 0;
}