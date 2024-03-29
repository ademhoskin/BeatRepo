package repo.beat.middleware.service;

public interface ZipFileS3Service {
    void uploadZipFile(String bucketName, String keyName, String filePath);
    void downloadZipFile(String bucketName, String keyName, String filePath);
    void deleteZipFile(String bucketName, String keyName);
}
