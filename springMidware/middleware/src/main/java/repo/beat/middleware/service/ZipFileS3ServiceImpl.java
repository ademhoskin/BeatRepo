package repo.beat.middleware.service;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class ZipFileS3ServiceImpl implements ZipFileS3Service {

    @Autowired
    private RestTemplate restTemplate;

    @Override
    public ResponseEntity<String> uploadZipFile(String bucketName, String objectKey, String filePath) {
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.MULTIPART_FORM_DATA);

        HashMap<String, Object> body = new HashMap<>();
        body.put("bucketName", bucketName);
        body.put("objectKey", objectKey);
        body.put("filePath", filePath);
        HttpEntity<Map<String, Object>> requestEntity = new HttpEntity<>(body, headers);

        ResponseEntity<String> response = restTemplate.exchange("http://localhost:8080/s3/upload", HttpMethod.POST, requestEntity, String.class);
        return response;

    }

    @Override
    public ResponseEntity<String> downloadZipFile(String bucketName, String objectKey, String filePath) {
        String url = "http://localhost:8080/s3/download?bucketName=" + bucketName + "&objectKey=" + objectKey + "&filePath=" + filePath;

        HttpHeaders headers = new HttpHeaders();
        HttpEntity<String> requestEntity = new HttpEntity<>(headers);

        ResponseEntity<String> response = restTemplate.exchange(url, HttpMethod.GET, requestEntity, String.class);
        return response;

    }

    @Override
    public ResponseEntity<String> deleteZipFile(String bucketName, String keyName) {
        String url = "http://localhost:8080/s3/delete?bucketName=" + bucketName + "&keyName=" + keyName;

        HttpHeaders headers = new HttpHeaders();
        HttpEntity<String> requestEntity = new HttpEntity<>(headers);

        ResponseEntity<String> response = restTemplate.exchange(url, HttpMethod.DELETE, requestEntity, String.class);
        return response;
        
    }
}