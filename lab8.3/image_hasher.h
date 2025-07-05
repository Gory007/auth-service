#ifndef IMAGE_HASHER_H
#define IMAGE_HASHER_H

#include <string>
#include <opencv2/opencv.hpp>

class ImageHasher {
public:
    static std::string calculateSHA256(const cv::Mat& image);
    static void modifyPixel(cv::Mat& image);
    static void saveResultsToFile(const std::string& originalHash, 
                                const std::string& modifiedHash, 
                                const std::string& filename);
};

#endif